package repositories

import (
	"context"
	"xrouting/internal/adapters/db"
	"xrouting/internal/domain/entities"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type ComponentRepository interface {
	CreateComponent(ctx context.Context, tableName string, component *entities.Component) error
	GetComponent(ctx context.Context, tableName string, componentID string) (*entities.Component, error)
	DeleteComponent(ctx context.Context, tableName string, componentID string) error
	ListComponents(ctx context.Context, tableName string) ([]*entities.Component, error)
}

type componentRepository struct {
	dynamo *db.DynamoDBClient
}

func NewComponentRepository(dynamo *db.DynamoDBClient) ComponentRepository {
	return &componentRepository{dynamo: dynamo}
}

func (r *componentRepository) CreateComponent(ctx context.Context, tableName string, component *entities.Component) error {
	return r.dynamo.PutItem(ctx, tableName, component)
}

func (r *componentRepository) GetComponent(ctx context.Context, tableName string, componentID string) (*entities.Component, error) {
	component := &entities.Component{}

	err := r.dynamo.GetItem(ctx, tableName, map[string]types.AttributeValue{
		"componentId": &types.AttributeValueMemberS{Value: componentID},
	}, component)

	if err != nil {
		return nil, err
	}

	return component, nil
}

func (r *componentRepository) DeleteComponent(ctx context.Context, tableName string, componentID string) error {
	return r.dynamo.DeleteItem(ctx, tableName, map[string]types.AttributeValue{
		"componentId": &types.AttributeValueMemberS{Value: componentID},
	})
}

func (r *componentRepository) ListComponents(ctx context.Context, tableName string) ([]*entities.Component, error) {
	components := make([]*entities.Component, 0)
	items, err := r.dynamo.ScanTable(ctx, tableName)
	if err != nil {
		return nil, err
	}

	for _, item := range items {
		component := &entities.Component{}
		err = component.Unmarshal(item)
		if err != nil {
			return nil, err
		}
		components = append(components, component)
	}
	return components, nil
}
