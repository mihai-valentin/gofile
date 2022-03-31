package handler

import (
	"github.com/gin-gonic/gin"
	"gofile/pkg/request"
	"gofile/pkg/response"
	"net/http"
)

func (h *FileHandler) uploadPrivateFile(c *gin.Context) {
	var privateFilesUploadRequest request.PrivateFilesUploadRequest

	if err := c.ShouldBind(&privateFilesUploadRequest); err != nil {
		response.Fail(c, response.NewErrorResponse(http.StatusBadRequest, err.Error()))
		return
	}

	files, err := h.service.UploadFiles(&FileUploadData{})
	if err != nil {
		response.Fail(c, response.NewErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusCreated, files)
}

func (h *FileHandler) getPrivateFile(c *gin.Context) {
	//uuid := c.Param("uuid")
	var privateFileAccessRequest request.PrivateFileAccessRequest

	if err := c.ShouldBind(&privateFileAccessRequest); err != nil {
		response.Fail(c, response.NewErrorResponse(http.StatusBadRequest, err.Error()))
		return
	}

	file, err := h.service.GetFileByFilter(&EntityFilter{})
	if err != nil {
		response.Fail(c, response.NewErrorResponse(http.StatusNotFound, err.Error()))
		return
	}

	c.File(file.Path)
}

func (h *FileHandler) deletePrivateFile(c *gin.Context) {
	//uuid := c.Param("uuid")
	var privateFileAccessRequest request.PrivateFileAccessRequest

	if err := c.ShouldBind(&privateFileAccessRequest); err != nil {
		response.Fail(c, response.NewErrorResponse(http.StatusBadRequest, err.Error()))
		return
	}

	err := h.service.DeleteFileByFilter(&EntityFilter{})
	if err != nil {
		response.Fail(c, response.NewErrorResponse(http.StatusNotFound, err.Error()))
		return
	}

	response.Success(c, response.NewCodeResponse(http.StatusNoContent))
}
