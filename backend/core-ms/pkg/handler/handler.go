package handler

import (
	"fullstack/backend/core-ms/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	authService *service.AuthService
}

func NewHandler(authService *service.AuthService) *Handler {
	return &Handler{authService: authService}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/core", h.userIdentity)
	{
		auth.GET("/test", h.test)
	}

	return router
}
