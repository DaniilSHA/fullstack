package handler

import (
	"fullstack/backend/auth-ms/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
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
