package api

import (
	"net/http"
	"xrouting/internal/domain/entities"
	"xrouting/internal/ports"

	"github.com/labstack/echo/v4"
)

type AnalyticHandler struct {
	analyticService ports.AnalyticService
}

func NewAnalyticHandler(analyticService ports.AnalyticService) *AnalyticHandler {
	return &AnalyticHandler{analyticService: analyticService}
}

func (h *AnalyticHandler) RegisterAnalytic(c echo.Context) error {
	analytic := new(entities.Analytic)
	if err := c.Bind(analytic); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	uid := GenerateRandomUID()
	analytic.ID = uid

	tableName := "analytics"

	err := h.analyticService.RegisterAnalytic(c.Request().Context(), tableName, analytic)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, analytic)
}

func (h *AnalyticHandler) FindAnalytic(c echo.Context) error {
	analyticID := c.Param("id")
	tableName := "analytics"

	analytic, err := h.analyticService.FindAnalytic(c.Request().Context(), tableName, analyticID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Analytic not found"})
	}

	return c.JSON(http.StatusOK, analytic)
}

func (h *AnalyticHandler) RemoveAnalytic(c echo.Context) error {
	analyticID := c.Param("id")
	tableName := "analytics"

	err := h.analyticService.RemoveAnalytic(c.Request().Context(), tableName, analyticID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Analytic deleted successfully"})
}

func (h *AnalyticHandler) Analytics(c echo.Context) error {
	tableName := "analytics"

	analytics, err := h.analyticService.Analytics(c.Request().Context(), tableName)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, analytics)
}
