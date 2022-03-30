package handler

import (
	"github.com/gin-gonic/gin"
	"gofile/pkg/entity"
	"net/http"
)

func (h *Handler) uploadPublicFile(c *gin.Context) {
	var filesUploadForm entity.FilesUploadForm

	if err := c.ShouldBind(&filesUploadForm); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	files, err := h.services.PublicFileManager.UploadFiles(filesUploadForm)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusCreated, files)
}

func (h *Handler) getPublicFile(c *gin.Context) {
	uuid := c.Param("uuid")

	file, err := h.services.PublicFileManager.GetFile(uuid)
	if err != nil {
		c.AbortWithStatusJSON(err.GetCode(), gin.H{"error": err.Error()})
		return
	}

	c.File(file.Path)
}

func (h *Handler) deletePublicFile(c *gin.Context) {
	uuid := c.Param("uuid")

	if err := h.services.PublicFileManager.DeleteFile(uuid); err != nil {
		c.AbortWithStatusJSON(err.GetCode(), gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusNoContent, gin.H{})
}
