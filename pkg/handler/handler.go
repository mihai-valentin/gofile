package handler

import (
	"github.com/gin-gonic/gin"
	"gofile/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRouter() *gin.Engine {
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
