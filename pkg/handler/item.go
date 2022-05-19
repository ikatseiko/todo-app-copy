package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/ikatseiko/todo-app-copy"
	"net/http"
	"strconv"
)

// @Summary CreateItem
// @Security ApiKeyAuth
// @Tags items
// @Description create todo item
// @ID create-item
// @Accept json
// @Produce json
// @Param id path int true "list id"
// @Param input body todo.TodoItem true "item info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/lists/{id}/items [post]
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

type getAllItemsResponse struct {
	Data []todo.TodoItem `json:"data"`
}

// @Summary GetAllItems
// @Security ApiKeyAuth
// @Tags items
// @Description get all todo items
// @ID get-all-items
// @Accept json
// @Produce json
// @Param id path int true "list id"
// @Success 200 {object} getAllItemsResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/lists/{id}/items [get]
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
	ctx.JSON(http.StatusOK, getAllItemsResponse{
		Data: items,
	})
}

// @Summary GetItemByID
// @Security ApiKeyAuth
// @Tags items
// @Description get todo Item
// @ID get-item-by-id
// @Accept json
// @Produce json
// @Param item_id path int true "item id"
// @Success 200 {object} todo.TodoItem
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/items/{item_id} [get]
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

// @Summary UpdateItem
// @Security ApiKeyAuth
// @Tags items
// @Description update todo Item
// @ID update-item
// @Accept json
// @Produce json
// @Param item_id path int true "item id"
// @Param input body todo.UpdateItemInput true "item info"
// @Success 200 {object} statusResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/items/{item_id} [put]
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

// @Summary DeleteItem
// @Security ApiKeyAuth
// @Tags items
// @Description delete todo Item
// @ID delete-item
// @Produce json
// @Param id path int true "item id"
// @Success 200 {object} statusResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/items/{id} [delete]
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
