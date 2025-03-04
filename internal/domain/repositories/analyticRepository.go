package repositories

import (
	"context"
	"xrouting/internal/adapters/db"
	"xrouting/internal/domain/entities"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type AnalyticRepository interface {
	CreateAnalytic(ctx context.Context, tableName string, analytic *entities.Analytic) error
	GetAnalytic(ctx context.Context, tableName string, analyticID string) (*entities.Analytic, error)
	DeleteAnalytic(ctx context.Context, tableName string, analyticID string) error
	ListAnalytics(ctx context.Context, tableName string) ([]*entities.Analytic, error)
}

type analyticRepository struct {
	dynamo *db.DynamoDBClient
}

func NewAnalyticRepository(dynamo *db.DynamoDBClient) AnalyticRepository {
	return &analyticRepository{dynamo: dynamo}
}

func (r *analyticRepository) CreateAnalytic(ctx context.Context, tableName string, analytic *entities.Analytic) error {
	return r.dynamo.PutItem(ctx, tableName, analytic)
}

func (r *analyticRepository) GetAnalytic(ctx context.Context, tableName string, analyticID string) (*entities.Analytic, error) {
	analytic := &entities.Analytic{}

	err := r.dynamo.GetItem(ctx, tableName, map[string]types.AttributeValue{
		"id": &types.AttributeValueMemberS{Value: analyticID},
	}, analytic)

	if err != nil {
		return nil, err
	}

	return analytic, nil
}

func (r *analyticRepository) DeleteAnalytic(ctx context.Context, tableName string, analyticID string) error {
	return r.dynamo.DeleteItem(ctx, tableName, map[string]types.AttributeValue{
		"id": &types.AttributeValueMemberS{Value: analyticID},
	})
}

func (r *analyticRepository) ListAnalytics(ctx context.Context, tableName string) ([]*entities.Analytic, error) {
	analytics := make([]*entities.Analytic, 0)
	items, err := r.dynamo.ScanTable(ctx, tableName)
	if err != nil {
		return nil, err
	}

	for _, item := range items {
		analytic := &entities.Analytic{}
		err = analytic.Unmarshal(item)
		if err != nil {
			return nil, err
		}
		analytics = append(analytics, analytic)
	}
	return analytics, nil
}
