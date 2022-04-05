package handler

import (
	"github.com/gin-gonic/gin"
	"gofile/pkg/contracts"
)

type FileHandler struct {
	service contracts.FileServiceInterface
}

func New(service contracts.FileServiceInterface) *FileHandler {
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
