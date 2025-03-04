package api

import (
	"net/http"
	"xrouting/internal/domain/entities"
	"xrouting/internal/ports"

	"github.com/labstack/echo/v4"
)

type ComponentHandler struct {
	componentService ports.ComponentService
}

func NewComponentHandler(componentService ports.ComponentService) *ComponentHandler {
	return &ComponentHandler{componentService: componentService}
}

func (h *ComponentHandler) RegisterComponent(c echo.Context) error {
	component := new(entities.Component)
	if err := c.Bind(component); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	uid := GenerateRandomUID()
	component.ID = uid

	tableName := "components"

	err := h.componentService.RegisterComponent(c.Request().Context(), tableName, component)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, component)
}

func (h *ComponentHandler) FindComponent(c echo.Context) error {
	componentID := c.Param("id")
	tableName := "components"

	component, err := h.componentService.FindComponent(c.Request().Context(), tableName, componentID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Component not found"})
	}

	return c.JSON(http.StatusOK, component)
}

func (h *ComponentHandler) RemoveComponent(c echo.Context) error {
	componentID := c.Param("id")
	tableName := "components"

	err := h.componentService.RemoveComponent(c.Request().Context(), tableName, componentID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Component removed successfully"})
}

func (h *ComponentHandler) Components(c echo.Context) error {
	tableName := "components"

	components, err := h.componentService.Components(c.Request().Context(), tableName)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, components)
}
