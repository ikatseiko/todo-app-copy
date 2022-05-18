package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/ikatseiko/todo-app-copy"
	"net/http"
	"strconv"
)

func (h *Handler) createItem(ctx *gin.Context) {
	userID, err := getUserID(ctx)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, "createItem | "+err.Error())
		return
	}

	listID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "createItem | invalid list id param")
		return
	}

	var input todo.TodoItem
	if err := ctx.BindJSON(&input); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "createItem | "+err.Error())
		return
	}

	id, err := h.services.TodoItem.Create(userID, listID, input)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, "createItem | "+err.Error())
		return
	}
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getAllItems(ctx *gin.Context) {
	userID, err := getUserID(ctx)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, "getAllItems | "+err.Error())
		return
	}

	listID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "getAllItems | invalid list id param")
		return
	}

	items, err := h.services.TodoItem.GetAll(userID, listID)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, "getAllItems | "+err.Error())
		return
	}
	ctx.JSON(http.StatusOK, items)
}

func (h *Handler) getItemByID(ctx *gin.Context) {
	userID, err := getUserID(ctx)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, "getItemByID 1 | "+err.Error())
		return
	}

	itemID, err := strconv.Atoi(ctx.Param("item_id"))
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "getItemByID 2 | invalid id param")
		return
	}
	list, err := h.services.TodoItem.GetByID(userID, itemID)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, "getItemByID 3 | "+err.Error())
		return
	}
	ctx.JSON(http.StatusOK, list)
}

func (h *Handler) updateItem(ctx *gin.Context) {

	userID, err := getUserID(ctx)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, "updateItem 1 | "+err.Error())
		return
	}

	id, err := strconv.Atoi(ctx.Param("item_id"))
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "updateItem 2 | invalid id param")
		return
	}

	var input todo.UpdateItemInput
	if err := ctx.BindJSON(&input); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "updateItem 3 | "+err.Error())
		return
	}

	if err := h.services.TodoItem.Update(userID, id, input); err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, "updateItem 4 | "+err.Error())
		return
	}
	ctx.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})

}

func (h *Handler) deleteItem(ctx *gin.Context) {
	userID, err := getUserID(ctx)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, "getItemByID 1 | "+err.Error())
		return
	}

	itemID, err := strconv.Atoi(ctx.Param("item_id"))
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "getItemByID 2 | invalid id param")
		return
	}

	if err := h.services.TodoItem.Delete(userID, itemID); err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, "getItemByID 3 | "+err.Error())
		return
	}
	ctx.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})

}
