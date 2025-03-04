package api

import (
	"net/http"
	"xrouting/internal/domain/entities"
	"xrouting/internal/ports"

	"github.com/labstack/echo/v4"
)

type MaintenanceHandler struct {
	maintenanceService ports.MaintenanceService
}

func NewMaintenanceHandler(maintenanceService ports.MaintenanceService) *MaintenanceHandler {
	return &MaintenanceHandler{maintenanceService: maintenanceService}
}

func (h *MaintenanceHandler) RegisterMaintenance(c echo.Context) error {
	maintenance := new(entities.Maintenance)
	if err := c.Bind(maintenance); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	uid := GenerateRandomUID()
	maintenance.ID = uid

	tableName := "maintenances"

	err := h.maintenanceService.RegisterMaintenance(c.Request().Context(), tableName, maintenance)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, maintenance)
}

func (h *MaintenanceHandler) FindMaintenance(c echo.Context) error {
	maintenanceID := c.Param("id")
	tableName := "maintenances"

	maintenance, err := h.maintenanceService.FindMaintenance(c.Request().Context(), tableName, maintenanceID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Maintenance not found"})
	}

	return c.JSON(http.StatusOK, maintenance)
}

func (h *MaintenanceHandler) RemoveMaintenance(c echo.Context) error {
	maintenanceID := c.Param("id")
	tableName := "maintenances"

	err := h.maintenanceService.RemoveMaintenance(c.Request().Context(), tableName, maintenanceID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Maintenance deleted successfully"})
}

func (h *MaintenanceHandler) Maintenances(c echo.Context) error {
	tableName := "maintenances"

	maintenances, err := h.maintenanceService.Maintenances(c.Request().Context(), tableName)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, maintenances)
}
