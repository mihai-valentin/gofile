package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const token = "123456789"

func (h *FileHandler) authorization(c *gin.Context) {
	authorizationHeader := c.GetHeader("Authorization")
	if authorizationHeader == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "authorization header missing"})
		return
	}

	authorizationHeaderParts := strings.Split(authorizationHeader, " ")
	if len(authorizationHeaderParts) != 2 || authorizationHeaderParts[0] != "Bearer" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "malformed authorization header"})
		return
	}

	if len(authorizationHeaderParts[1]) == 0 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "malformed authorization header"})
		return
	}

	if authorizationHeaderParts[1] != token {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "wrong credentials"})
		return
	}
}
