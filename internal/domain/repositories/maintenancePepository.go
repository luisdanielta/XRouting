package repositories

import (
	"context"
	"xrouting/internal/adapters/db"
	"xrouting/internal/domain/entities"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type MaintenanceRepository interface {
	CreateMaintenance(ctx context.Context, tableName string, maintenance *entities.Maintenance) error
	GetMaintenance(ctx context.Context, tableName string, maintenanceID string) (*entities.Maintenance, error)
	DeleteMaintenance(ctx context.Context, tableName string, maintenanceID string) error
	ListMaintenances(ctx context.Context, tableName string) ([]*entities.Maintenance, error)
}

type maintenanceRepository struct {
	dynamo *db.DynamoDBClient
}

func NewMaintenanceRepository(dynamo *db.DynamoDBClient) MaintenanceRepository {
	return &maintenanceRepository{dynamo: dynamo}
}

func (r *maintenanceRepository) CreateMaintenance(ctx context.Context, tableName string, maintenance *entities.Maintenance) error {
	return r.dynamo.PutItem(ctx, tableName, maintenance)
}

func (r *maintenanceRepository) GetMaintenance(ctx context.Context, tableName string, maintenanceID string) (*entities.Maintenance, error) {
	maintenance := &entities.Maintenance{}

	err := r.dynamo.GetItem(ctx, tableName, map[string]types.AttributeValue{
		"id": &types.AttributeValueMemberS{Value: maintenanceID},
	}, maintenance)

	if err != nil {
		return nil, err
	}

	return maintenance, nil
}

func (r *maintenanceRepository) DeleteMaintenance(ctx context.Context, tableName string, maintenanceID string) error {
	return r.dynamo.DeleteItem(ctx, tableName, map[string]types.AttributeValue{
		"id": &types.AttributeValueMemberS{Value: maintenanceID},
	})
}

func (r *maintenanceRepository) ListMaintenances(ctx context.Context, tableName string) ([]*entities.Maintenance, error) {
	maintenances := make([]*entities.Maintenance, 0)
	items, err := r.dynamo.ScanTable(ctx, tableName)
	if err != nil {
		return nil, err
	}

	for _, item := range items {
		maintenance := &entities.Maintenance{}
		err = maintenance.Unmarshal(item)
		if err != nil {
			return nil, err
		}
		maintenances = append(maintenances, maintenance)
	}
	return maintenances, nil
}
