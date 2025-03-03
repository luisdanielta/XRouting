package db

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type DynamoDBEntity interface {
	GetID() string
}

type DynamoDBMapper[T DynamoDBEntity] struct{}

// Utiliza attributevalue.MarshalMap para convertir la entidad a un map compatible con DynamoDB
func (m *DynamoDBMapper[T]) ToDynamoDBMap(entity T) (map[string]types.AttributeValue, error) {
	return attributevalue.MarshalMap(entity)
}

// Utiliza attributevalue.UnmarshalMap para convertir el map de DynamoDB a la entidad
func (m *DynamoDBMapper[T]) FromDynamoDBMap(dynamoMap map[string]types.AttributeValue) (T, error) {
	var entity T
	err := attributevalue.UnmarshalMap(dynamoMap, &entity)
	return entity, err
}

func (m *DynamoDBMapper[T]) ToDynamoDBKeyID(id string) map[string]types.AttributeValue {
	return map[string]types.AttributeValue{
		"ID": &types.AttributeValueMemberS{Value: id},
	}
}

type DynamoDBRepository interface {
	PutItem(ctx context.Context, table string, item map[string]types.AttributeValue) error
	GetItem(ctx context.Context, table string, key map[string]types.AttributeValue) (map[string]types.AttributeValue, error)
	UpdateItem(ctx context.Context, table string, key map[string]types.AttributeValue, updateExpression string, expressionAttributeValues map[string]types.AttributeValue) error
	DeleteItem(ctx context.Context, table string, key map[string]types.AttributeValue) error
}

type DynamoDBClient struct {
	client *dynamodb.Client
}

func NewDynamoDBClient(cfg aws.Config) *DynamoDBClient {
	client := dynamodb.NewFromConfig(cfg)
	return &DynamoDBClient{client: client}
}

func (db *DynamoDBClient) PutItem(ctx context.Context, table string, item map[string]types.AttributeValue) error {
	input := &dynamodb.PutItemInput{
		TableName: &table,
		Item:      item,
	}
	_, err := db.client.PutItem(ctx, input)
	if err != nil {
		return fmt.Errorf("error al insertar item: %w", err)
	}
	return nil
}

func (db *DynamoDBClient) GetItem(ctx context.Context, table string, key map[string]types.AttributeValue) (map[string]types.AttributeValue, error) {
	input := &dynamodb.GetItemInput{
		TableName: &table,
		Key:       key,
	}
	result, err := db.client.GetItem(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("error al obtener item: %w", err)
	}
	return result.Item, nil
}

func (db *DynamoDBClient) UpdateItem(ctx context.Context, table string, key map[string]types.AttributeValue, updateExpression string, expressionAttributeValues map[string]types.AttributeValue) error {
	input := &dynamodb.UpdateItemInput{
		TableName:                 &table,
		Key:                       key,
		UpdateExpression:          &updateExpression,
		ExpressionAttributeValues: expressionAttributeValues,
	}
	_, err := db.client.UpdateItem(ctx, input)
	if err != nil {
		return fmt.Errorf("error al actualizar item: %w", err)
	}
	return nil
}

func (db *DynamoDBClient) DeleteItem(ctx context.Context, table string, key map[string]types.AttributeValue) error {
	input := &dynamodb.DeleteItemInput{
		TableName: &table,
		Key:       key,
	}
	_, err := db.client.DeleteItem(ctx, input)
	if err != nil {
		return fmt.Errorf("error al eliminar item: %w", err)
	}
	return nil
}

type ItemService struct {
	repository DynamoDBRepository
}

func NewItemService(repo DynamoDBRepository) *ItemService {
	return &ItemService{
		repository: repo,
	}
}

func (s *ItemService) CreateItem(ctx context.Context, table string, item map[string]types.AttributeValue) error {
	return s.repository.PutItem(ctx, table, item)
}

func (s *ItemService) ReadItem(ctx context.Context, table string, key map[string]types.AttributeValue) (map[string]types.AttributeValue, error) {
	return s.repository.GetItem(ctx, table, key)
}

func (s *ItemService) UpdateItem(ctx context.Context, table string, key map[string]types.AttributeValue, updateExpression string, expressionAttributeValues map[string]types.AttributeValue) error {
	return s.repository.UpdateItem(ctx, table, key, updateExpression, expressionAttributeValues)
}

func (s *ItemService) DeleteItem(ctx context.Context, table string, key map[string]types.AttributeValue) error {
	return s.repository.DeleteItem(ctx, table, key)
}

func (db *DynamoDBClient) GetDBSession() string {
	return db.client.Options().Region
}
