package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/ikatseiko/todo-app-copy"
	"net/http"
	"strconv"
)

// @Summary CreateList
// @Security ApiKeyAuth
// @Tags lists
// @Description create todo list
// @ID create-list
// @Accept json
// @Produce json
// @Param input body todo.TodoList true "list info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/lists [post]
func (h *Handler) createList(ctx *gin.Context) {
	userID, err := getUserID(ctx)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, "createList | "+err.Error())
		return
	}
	var input todo.TodoList
	if err := ctx.BindJSON(&input); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "createList | "+err.Error())
		return
	}
	id, err := h.services.TodoList.Create(userID, input)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, "createList | "+err.Error())
		return
	}
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type getAllListsResponse struct {
	Data []todo.TodoList `json:"data"`
}

// @Summary GetAllLists
// @Security ApiKeyAuth
// @Tags lists
// @Description get all todo lists
// @ID get-all-lists
// @Accept json
// @Produce json
// @Success 200 {object} getAllListsResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/lists [get]
func (h *Handler) getAllLists(ctx *gin.Context) {
	userID, err := getUserID(ctx)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, "getAllLists 1 | "+err.Error())
		return
	}
	lists, err := h.services.TodoList.GetAll(userID)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, "getAllLists 2 | "+err.Error())
		return
	}
	ctx.JSON(http.StatusOK, getAllListsResponse{
		Data: lists,
	})
}

// @Summary GetListByID
// @Security ApiKeyAuth
// @Tags lists
// @Description get todo list
// @ID get-list-by-id
// @Accept json
// @Produce json
// @Param id path int true "list id"
// @Success 200 {object} todo.TodoList
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/lists/{id} [get]
func (h *Handler) getListByID(ctx *gin.Context) {
	userID, err := getUserID(ctx)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, "getListByID 1 | "+err.Error())
		return
	}

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "getListByID 2 | invalid id param")
		return
	}
	list, err := h.services.TodoList.GetByID(userID, id)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, "getListByID 3 | "+err.Error())
		return
	}
	ctx.JSON(http.StatusOK, list)
}

// @Summary UpdateList
// @Security ApiKeyAuth
// @Tags lists
// @Description update todo list
// @ID update-list
// @Accept json
// @Produce json
// @Param id path int true "list id"
// @Param input body todo.UpdateListInput true "list info"
// @Success 200 {object} statusResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/lists/{id} [put]
func (h *Handler) updateList(ctx *gin.Context) {
	userID, err := getUserID(ctx)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, "updateList 1 | "+err.Error())
		return
	}

	listID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "updateList 2 | invalid id param")
		return
	}

	var input todo.UpdateListInput
	if err := ctx.BindJSON(&input); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "updateList 3 | "+err.Error())
		return
	}

	if err := h.services.TodoList.Update(userID, listID, input); err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, "updateList 4 | "+err.Error())
		return
	}
	ctx.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

// @Summary DeleteList
// @Security ApiKeyAuth
// @Tags lists
// @Description delete todo list
// @ID delete-list
// @Produce json
// @Param id path int true "list id"
// @Success 200 {object} statusResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/lists/{id} [delete]
func (h *Handler) deleteList(ctx *gin.Context) {
	userID, err := getUserID(ctx)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, "deleteList 1 | "+err.Error())
		return
	}

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "deleteList 2 | invalid id param")
		return
	}
	if err := h.services.TodoList.Delete(userID, id); err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, "deleteList 3 | "+err.Error())
		return
	}
	ctx.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
