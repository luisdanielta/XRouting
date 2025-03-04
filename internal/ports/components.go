package ports

import (
	"context"
	"xrouting/internal/domain/entities"
	"xrouting/internal/domain/repositories"
)

type ComponentService interface {
	RegisterComponent(ctx context.Context, tableName string, component *entities.Component) error
	FindComponent(ctx context.Context, tableName string, componentID string) (*entities.Component, error)
	RemoveComponent(ctx context.Context, tableName string, componentID string) error
	Components(ctx context.Context, tableName string) ([]*entities.Component, error)
}

type componentService struct {
	componentRepo repositories.ComponentRepository
}

func NewComponentService(componentRepo repositories.ComponentRepository) ComponentService {
	return &componentService{componentRepo: componentRepo}
}

func (s *componentService) RegisterComponent(ctx context.Context, tableName string, component *entities.Component) error {
	return s.componentRepo.CreateComponent(ctx, tableName, component)
}

func (s *componentService) FindComponent(ctx context.Context, tableName string, componentID string) (*entities.Component, error) {
	return s.componentRepo.GetComponent(ctx, tableName, componentID)
}

func (s *componentService) RemoveComponent(ctx context.Context, tableName string, componentID string) error {
	return s.componentRepo.DeleteComponent(ctx, tableName, componentID)
}

func (s *componentService) Components(ctx context.Context, tableName string) ([]*entities.Component, error) {
	return s.componentRepo.ListComponents(ctx, tableName)
}
