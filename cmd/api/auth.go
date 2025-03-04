package api

import (
	"fmt"
	"net/http"
	"xrouting/internal/domain/entities"
	"xrouting/internal/ports"

	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	authService ports.AuthService
}

type UserDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func NewAuthHandler(authService ports.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

func (h *AuthHandler) SignUp(c echo.Context) error {

	user := new(entities.User)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})

	}

	uid := GenerateRandomUID()
	user.ID = uid

	err := h.authService.SignUp(c.Request().Context(), user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, map[string]string{"message": "User registered successfully"})
}

func (h *AuthHandler) SignIn(c echo.Context) error {
	credentials := new(UserDTO)
	fmt.Println(credentials)
	if err := c.Bind(credentials); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	token, err := h.authService.SignIn(c.Request().Context(), credentials.Username, credentials.Password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid credentials"})
	}

	return c.JSON(http.StatusOK, map[string]string{"token": token})
}
