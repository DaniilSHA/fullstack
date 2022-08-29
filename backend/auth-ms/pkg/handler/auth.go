package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type signInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) registration(c *gin.Context) {
	fmt.Print("REGISTRATION")
}

func (h *Handler) login(c *gin.Context) {
	fmt.Print("LOGIN")

}

func (h *Handler) refresh(c *gin.Context) {
	fmt.Print("REFRESH")

}
