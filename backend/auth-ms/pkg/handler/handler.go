package handler

import (
	"fullstack/backend/auth-ms/pkg/service"
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

	auth := router.Group("/auth")
	{
		auth.POST("/registration", h.registration)
		auth.POST("/login", h.login)
		auth.POST("/refresh", h.refresh)
	}

	return router
}
