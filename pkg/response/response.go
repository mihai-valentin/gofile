package response

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type JsonResponseInterface interface {
	GetCode() int
	GetText() string
}

func Fail(c *gin.Context, response JsonResponseInterface) {
	logrus.Error(response.GetText())
	c.AbortWithStatusJSON(response.GetCode(), response.GetText())
}

func Success(c *gin.Context, response JsonResponseInterface) {
	logrus.Error(response.GetText())
	c.JSON(response.GetCode(), response.GetText())
}
