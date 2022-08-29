package handler

import (
	"fmt"
	"fullstack/backend/auth-ms/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

//type signInInput struct {
//	Username string `json:"username" binding:"required"`
//	Password string `json:"password" binding:"required"`
//}

func (h *Handler) registration(c *gin.Context) {
	var input models.UserDto

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	fmt.Print("REGISTRATION")
}

func (h *Handler) login(c *gin.Context) {
	fmt.Print("LOGIN")

}

func (h *Handler) refresh(c *gin.Context) {
	fmt.Print("REFRESH")

}
