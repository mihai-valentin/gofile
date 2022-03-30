package handler

import (
	"github.com/gin-gonic/gin"
	"gofile/pkg/request"
	"gofile/pkg/response"
	"net/http"
)

func (h *Handler) uploadPublicFile(c *gin.Context) {
	var pubicFilesUploadRequest request.PubicFilesUploadRequest

	if err := c.ShouldBind(&pubicFilesUploadRequest); err != nil {
		response.Fail(c, response.NewErrorResponse(http.StatusBadRequest, err.Error()))
		return
	}

	files, err := h.services.PublicFileManager.UploadFiles(pubicFilesUploadRequest)
	if err != nil {
		response.Fail(c, response.NewErrorResponse(http.StatusInternalServerError, err.Error()))
		return
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
