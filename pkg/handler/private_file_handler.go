package handler

import (
	"github.com/gin-gonic/gin"
	"gofile/pkg/entity"
	"net/http"
)

func (h *Handler) uploadPrivateFile(c *gin.Context) {
	var filesUploadForm entity.FilesUploadForm

	if err := c.ShouldBind(&filesUploadForm); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	files, err := h.services.PrivateFileManager.UploadFiles(filesUploadForm)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusCreated, files)
}

func (h *Handler) getPrivateFile(c *gin.Context) {
	uuid := c.Param("uuid")

	file, err := h.services.PrivateFileManager.GetFile(uuid)
	if err != nil {
		c.AbortWithStatusJSON(err.GetCode(), gin.H{"error": err.Error()})
		return
	}

	if !h.services.PrivateFileManager.MatchOwner(file, getFileOwner(c)) {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{})
		return
	}

	c.File(file.Path)
}

func (h *Handler) deletePrivateFile(c *gin.Context) {
	uuid := c.Param("uuid")

	file, err := h.services.PrivateFileManager.GetFile(uuid)
	if err != nil {
		c.AbortWithStatusJSON(err.GetCode(), gin.H{"error": err.Error()})
	}

	if !h.services.PrivateFileManager.MatchOwner(file, getFileOwner(c)) {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{})
	}

	if err := h.services.PrivateFileManager.DeleteFile(uuid); err != nil {
		c.AbortWithStatusJSON(err.GetCode(), gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusNoContent, gin.H{})
}
