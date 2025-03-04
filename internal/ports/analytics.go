package ports

import (
	"context"
	"xrouting/internal/domain/entities"
	"xrouting/internal/domain/repositories"
)

type AnalyticService interface {
	RegisterAnalytic(ctx context.Context, tableName string, analytic *entities.Analytic) error
	FindAnalytic(ctx context.Context, tableName string, analyticID string) (*entities.Analytic, error)
	RemoveAnalytic(ctx context.Context, tableName string, analyticID string) error
	Analytics(ctx context.Context, tableName string) ([]*entities.Analytic, error)
}

type analyticService struct {
	analyticRepo repositories.AnalyticRepository
}

func NewAnalyticService(analyticRepo repositories.AnalyticRepository) AnalyticService {
	return &analyticService{analyticRepo: analyticRepo}
}

func (s *analyticService) RegisterAnalytic(ctx context.Context, tableName string, analytic *entities.Analytic) error {
	return s.analyticRepo.CreateAnalytic(ctx, tableName, analytic)
}

func (s *analyticService) FindAnalytic(ctx context.Context, tableName string, analyticID string) (*entities.Analytic, error) {
	return s.analyticRepo.GetAnalytic(ctx, tableName, analyticID)
}

func (s *analyticService) RemoveAnalytic(ctx context.Context, tableName string, analyticID string) error {
	return s.analyticRepo.DeleteAnalytic(ctx, tableName, analyticID)
}

func (s *analyticService) Analytics(ctx context.Context, tableName string) ([]*entities.Analytic, error) {
	return s.analyticRepo.ListAnalytics(ctx, tableName)
}
