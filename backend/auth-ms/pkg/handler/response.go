package handler

import (
	"fullstack/backend/auth-ms/models"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type error struct {
	Message string `json:"message"`
}

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	logrus.Error(message)
	c.AbortWithStatusJSON(statusCode, error{message})
}

func newOkResponseWithTokens(c *gin.Context, statusCode int, tokens *models.Tokens) {
	c.AbortWithStatusJSON(statusCode, tokens)
}

func newOkResponse(c *gin.Context, statusCode int) {
	c.AbortWithStatus(statusCode)
}
