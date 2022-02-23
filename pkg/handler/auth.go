package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/ikatseiko/todo-app-copy"
	"net/http"
)

func (h *Handler) signUp(ctx *gin.Context) {
	var input todo.User

	if err := ctx.BindJSON(&input); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.services.Autorization.CreateUser(input)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type signInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) signIn(ctx *gin.Context) {
	var input signInInput

	if err := ctx.BindJSON(&input); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "signIn | "+err.Error())
		return
	}
	token, err := h.services.Autorization.GenerateToken(input.Username, input.Password)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, "signIn(token) | "+err.Error())
		return
	}
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})

}
