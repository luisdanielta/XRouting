package ports

import (
	"context"
	"xrouting/internal/domain/entities"
	"xrouting/internal/domain/repositories"
)

type MaintenanceService interface {
	RegisterMaintenance(ctx context.Context, tableName string, maintenance *entities.Maintenance) error
	FindMaintenance(ctx context.Context, tableName string, maintenanceID string) (*entities.Maintenance, error)
	RemoveMaintenance(ctx context.Context, tableName string, maintenanceID string) error
	Maintenances(ctx context.Context, tableName string) ([]*entities.Maintenance, error)
}

type maintenanceService struct {
	maintenanceRepo repositories.MaintenanceRepository
}

func NewMaintenanceService(maintenanceRepo repositories.MaintenanceRepository) MaintenanceService {
	return &maintenanceService{maintenanceRepo: maintenanceRepo}
}

func (s *maintenanceService) RegisterMaintenance(ctx context.Context, tableName string, maintenance *entities.Maintenance) error {
	return s.maintenanceRepo.CreateMaintenance(ctx, tableName, maintenance)
}

func (s *maintenanceService) FindMaintenance(ctx context.Context, tableName string, maintenanceID string) (*entities.Maintenance, error) {
	return s.maintenanceRepo.GetMaintenance(ctx, tableName, maintenanceID)
}

func (s *maintenanceService) RemoveMaintenance(ctx context.Context, tableName string, maintenanceID string) error {
	return s.maintenanceRepo.DeleteMaintenance(ctx, tableName, maintenanceID)
}

func (s *maintenanceService) Maintenances(ctx context.Context, tableName string) ([]*entities.Maintenance, error) {
	return s.maintenanceRepo.ListMaintenances(ctx, tableName)
}
