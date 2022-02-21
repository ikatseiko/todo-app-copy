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
	}
}

func (h *Handler) signIn(ctx *gin.Context) {

}
