package api

import (
	"crypto/rand"
	"encoding/hex"
	"net/http"
	"xrouting/internal/domain/entities"
	"xrouting/internal/ports"

	"github.com/labstack/echo/v4"
)

func GenerateRandomUID() string {
	bytes := make([]byte, 16)
	_, err := rand.Read(bytes)
	if err != nil {
		panic(err)
	}
	return hex.EncodeToString(bytes)
}

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

	uid := GenerateRandomUID()
	user.ID = uid

	tableName := "users"

	err := h.userService.RegisterUser(c.Request().Context(), tableName, user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, user)
}

func (h *UserHandler) FindUser(c echo.Context) error {
	userID := c.Param("id")
	tableName := "users"

	user, err := h.userService.FindUser(c.Request().Context(), tableName, userID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "User not found"})
	}

	return c.JSON(http.StatusOK, user)
}

func (h *UserHandler) RemoveUser(c echo.Context) error {
	userID := c.Param("id")
	tableName := "users"

	err := h.userService.RemoveUser(c.Request().Context(), tableName, userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.NoContent(http.StatusNoContent)
}
