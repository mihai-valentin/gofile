package handler

import (
	"github.com/gin-gonic/gin"
	"gofile/pkg/data"
	"gofile/pkg/entity"
)

type Condition struct {
	Field    string
	Operator string
	Value    string
}

type EntityFilter struct {
	Conditions []*Condition
}

type FileServiceInterface interface {
	UploadFiles(filesUploadData []*data.UploadFileData) ([]*entity.File, error)
	GetFile(uuid string, ownerSign string) (string, error)
	DeleteFile(uuid string, ownerSign string) error
}

type FileHandler struct {
	service FileServiceInterface
}

func New(service FileServiceInterface) *FileHandler {
	return &FileHandler{service}
}

func (h *FileHandler) InitRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/health", h.checkHealth)

	api := router.Group("/api", h.authorization)
	{
		files := api.Group("/files")
		{
			files.POST("", h.uploadFile)
			files.GET("/:uuid", h.getFile)
			files.DELETE("/:uuid", h.deleteFile)
		}
	}

	return router
}
