package handler

import (
	"github.com/gin-gonic/gin"
	"gofile/pkg/entity"
	"net/http"
)

func (h *Handler) checkFileOwnerExistence(c *gin.Context) {
	ownerGuid, ok := c.GetPostForm("owner_guid")
	if !ok {
		c.AbortWithStatusJSON(
			http.StatusUnprocessableEntity,
			gin.H{"error": "File owner guid (`owner_guid`) field missing"},
		)
	}

	ownerType, ok := c.GetPostForm("owner_type")
	if !ok {
		c.AbortWithStatusJSON(
			http.StatusUnprocessableEntity,
			gin.H{"error": "File owner type (`owner_type`) field missing"},
		)
	}

	c.Set("owner_guid", ownerGuid)
	c.Set("owner_type", ownerType)
}

func getFileOwner(c *gin.Context) *entity.FileOwner {
	ownerGuid, _ := c.Get("owner_guid")
	ownerType, _ := c.Get("owner_type")

	return &entity.FileOwner{
		Guid: ownerGuid.(string),
		Type: ownerType.(string),
	}
}
