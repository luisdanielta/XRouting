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
