package handler

import (
	"github.com/gin-gonic/gin"
	_ "gofile/pkg/entity"
	gofile "gofile/pkg/entity"
	"net/http"
)

func (h *Handler) uploadPublicFile(c *gin.Context) {
	var publicFiles gofile.PublicFilesList

	if err := c.ShouldBind(&publicFiles); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	files, err := h.services.PublicFile.UploadFiles(publicFiles)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusCreated, files)
}

func (h *Handler) getPublicFile(c *gin.Context) {
	uuid := c.Param("uuid")

	file, err := h.services.PublicFile.GetFile(uuid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.File(file.Path)
}

func (h *Handler) deletePublicFile(c *gin.Context) {
	uuid := c.Param("uuid")
	err := h.services.PublicFile.DeleteFile(uuid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
}
