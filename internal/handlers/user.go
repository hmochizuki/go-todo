package handlers

import (
	"go-todo/internal/domain"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	Service domain.UserService
}

func NewUserHandler(service domain.UserService) *UserHandler {
	return &UserHandler{Service: service}
}

func (h *UserHandler) CreateUser(c echo.Context) error {
	var requestedUser domain.CreateUserRequest
	if err := c.Bind(&requestedUser); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}
	err := h.Service.CreateUser(requestedUser)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, nil)
}
