package handler

import (
	"fmt"
	"fullstack/backend/auth-ms/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) registration(c *gin.Context) {
	var input models.UserDto

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	userId, err := h.authService.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusConflict, err.Error())
		return
	}

	newOkResponse(c, http.StatusCreated, userId)
}

func (h *Handler) login(c *gin.Context) {
	var input models.UserDto

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	_, err := h.authService.CheckUserCredentials(input)
	if err != nil {
		newErrorResponse(c, http.StatusConflict, err.Error())
		return
	}

	newOkResponse(c, http.StatusAccepted, "user confirmed")
}

func (h *Handler) refresh(c *gin.Context) {
	fmt.Print("REFRESH")
}
