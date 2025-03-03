package api

import (
	"net/http"
	"xrouting/internal/domain/entities"
	"xrouting/internal/ports"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userService ports.UserService
}

func NewUserHandler(userService ports.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) RegisterUser(c echo.Context) error {

	user := new(entities.User)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	tableName := "UsersTable"

	err := h.userService.RegisterUser(c.Request().Context(), tableName, user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, user)
}

func (h *UserHandler) FindUser(c echo.Context) error {
	userID := c.Param("id")
	tableName := "UsersTable"

	user, err := h.userService.FindUser(c.Request().Context(), tableName, userID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "User not found"})
	}

	return c.JSON(http.StatusOK, user)
}

func (h *UserHandler) RemoveUser(c echo.Context) error {
	userID := c.Param("id")
	tableName := "UsersTable"

	err := h.userService.RemoveUser(c.Request().Context(), tableName, userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.NoContent(http.StatusNoContent)
}
