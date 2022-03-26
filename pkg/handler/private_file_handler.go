package handler

import (
	"github.com/gin-gonic/gin"
	gofile "gofile/pkg/entity"
	"net/http"
)

func (h *Handler) uploadPrivateFile(c *gin.Context) {
	var privateFiles gofile.PrivateFilesList

	if err := c.ShouldBind(&privateFiles); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	files, err := h.services.PrivateFile.UploadFiles(privateFiles)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusCreated, files)
}

func (h *Handler) getPrivateFile(c *gin.Context) {
	var fileOwner gofile.FileOwner

	if err := c.ShouldBind(&fileOwner); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File owner data missing"})
		return
	}

	uuid := c.Param("uuid")
	file, err := h.services.PrivateFile.GetFile(uuid, fileOwner)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.File(file.Path)
}

func (h *Handler) deletePrivateFile(c *gin.Context) {
	var fileOwner gofile.FileOwner

	if err := c.ShouldBind(&fileOwner); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File owner data missing"})
	}

	uuid := c.Param("uuid")
	err := h.services.PrivateFile.DeleteFile(uuid, fileOwner)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
}
