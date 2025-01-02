package handlers

import (
	"go-todo/internal/domain"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type TodoHandler struct {
	Service domain.TodoService
}

func NewTodoHandler(service domain.TodoService) *TodoHandler {
	return &TodoHandler{Service: service}
}

func (h *TodoHandler) GetAllTodos(c echo.Context) error {
	todos, err := h.Service.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, todos)
}

func (h *TodoHandler) CreateTodo(c echo.Context) error {
	var requestedTodo domain.CreateTodoRequest
	if err := c.Bind(&requestedTodo); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}
	err := h.Service.Create(requestedTodo)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, nil)
}

func (h *TodoHandler) UpdateTodoStatus(c echo.Context) error {
	var requestedTodo domain.UpdateTodoStatusRequest
	if err := c.Bind(&requestedTodo); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid ID"})
	}
	err = h.Service.UpdateTodoStatus(uint(id), requestedTodo)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, nil)
}

func (h *TodoHandler) DeleteTodo(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid ID"})
	}
	if err := h.Service.Delete(id); err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "todo not found"})
	}
	return c.NoContent(http.StatusNoContent)
}
