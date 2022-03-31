package handler

import (
	"github.com/gin-gonic/gin"
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

type FileUploadData struct {
}

type FileServiceInterface interface {
	UploadFiles(fileUploadData *FileUploadData) ([]*entity.File, error)
	GetFileByFilter(filter *EntityFilter) (*entity.File, error)
	DeleteFileByFilter(filter *EntityFilter) error
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
			public := files.Group("/public")
			{
				public.POST("", h.uploadPublicFile)
				public.GET("/:uuid", h.getPublicFile)
				public.DELETE("/:uuid", h.deletePublicFile)
			}

			private := files.Group("/private")
			{
				private.POST("", h.uploadPrivateFile)
				private.GET("/:uuid", h.getPrivateFile)
				private.DELETE("/:uuid", h.deletePrivateFile)
			}
		}
	}

	return router
}
