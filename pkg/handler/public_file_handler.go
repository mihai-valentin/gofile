package handler

import (
	"github.com/gin-gonic/gin"
	"gofile/pkg/request"
	"gofile/pkg/response"
	"net/http"
)

func (h *FileHandler) uploadPublicFile(c *gin.Context) {
	var pubicFilesUploadRequest request.PubicFilesUploadRequest

	if err := c.ShouldBind(&pubicFilesUploadRequest); err != nil {
		response.Fail(c, response.NewErrorResponse(http.StatusBadRequest, err.Error()))
		return
	}

	files, err := h.service.UploadFiles(pubicFilesUploadRequest)
	if err != nil {
		response.Fail(c, response.NewErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusCreated, files)
}

func (h *FileHandler) getPublicFile(c *gin.Context) {
	uuid := c.Param("uuid")

	file, err := h.service.GetFileByFilter(uuid)
	if err != nil {
		response.Fail(c, response.NewErrorResponse(http.StatusNotFound, err.Error()))
		return
	}

	c.File(file.Path)
}

func (h *FileHandler) deletePublicFile(c *gin.Context) {
	uuid := c.Param("uuid")

	if err := h.service.DeleteFileByFilter(uuid); err != nil {
		response.Fail(c, response.NewErrorResponse(http.StatusNotFound, err.Error()))
		return
	}

	response.Success(c, response.NewCodeResponse(http.StatusNoContent))
}
