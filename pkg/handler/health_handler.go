package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *FileHandler) checkHealth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"health": "OK"})
}
