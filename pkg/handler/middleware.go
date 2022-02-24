package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userID"
)

func (h *Handler) userIdentity(ctx *gin.Context) {
	header := ctx.GetHeader(authorizationHeader)
	if header == "" {
		newErrorResponse(ctx, http.StatusUnauthorized, "userIdentity | empty auth header")
		return
	}
	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		newErrorResponse(ctx, http.StatusUnauthorized, "userIdentity | invalid auth header")
		return
	}
	userID, err := h.services.Autorization.ParsToken(headerParts[1])
	if err != nil {
		newErrorResponse(ctx, http.StatusUnauthorized, "userIdentity | "+err.Error())
		return
	}
	ctx.Set(userCtx, userID)
	// pars token
}
