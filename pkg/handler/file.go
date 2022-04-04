package handler

import (
	"github.com/gin-gonic/gin"
	"gofile/pkg/mapper"
	"gofile/pkg/request"
	"net/http"
)

func (h *FileHandler) uploadFile(c *gin.Context) {
	var uploadRequest request.FilesUploadRequest
	dataMapper := new(mapper.UploadDataMapper)

	if err := c.ShouldBind(&uploadRequest); err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	filesUploadData := dataMapper.MapFromRequest(&uploadRequest)
	files, err := h.service.UploadFiles(filesUploadData)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusCreated, files)
}

func (h *FileHandler) getFile(c *gin.Context) {
	uuid := c.Param("uuid")

	filePath, err := h.service.GetFile(uuid, c.Query("owner_sign"))
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.File(filePath)
}

func (h *FileHandler) deleteFile(c *gin.Context) {
	uuid := c.Param("uuid")

	if err := h.service.DeleteFile(uuid, c.Query("owner_sign")); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.Status(http.StatusNoContent)
}
