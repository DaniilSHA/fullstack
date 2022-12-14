package handler

import (
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

	err := h.authService.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusConflict, err.Error())
		return
	}

	newOkResponse(c, http.StatusCreated)
}

func (h *Handler) login(c *gin.Context) {
	var input models.UserDto

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	tokens, err := h.authService.CheckUserCredentials(input)
	if err != nil {
		newErrorResponse(c, http.StatusConflict, err.Error())
		return
	}

	newOkResponseWithTokens(c, http.StatusOK, tokens)
}

func (h *Handler) refresh(c *gin.Context) {
	var input models.Tokens

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	tokens, err := h.authService.ValidateAndRefreshTokens(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	newOkResponseWithTokens(c, http.StatusOK, tokens)
}
