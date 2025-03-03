package db

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

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

// EnsureTableExists checks if a DynamoDB table with the specified name exists.
// If the table does not exist, it creates the table with a predefined schema.
// If the table already exists, it returns an error.
//
// Parameters:
//
//	ctx - The context for the operation.
//	tableName - The name of the DynamoDB table to check or create.
//
// Returns:
//
//	An error if the table already exists or if there was an error creating the table.
func (db *DynamoDBClient) EnsureTableExists(ctx context.Context, tableName string) error {
	_, err := db.client.DescribeTable(ctx, &dynamodb.DescribeTableInput{
		TableName: &tableName,
	})

	if err == nil {
		return fmt.Errorf("table '%s' already exists", tableName)
	}

	fmt.Printf("Making sure table '%s' exists...\n", tableName)
	_, err = db.client.CreateTable(ctx, &dynamodb.CreateTableInput{
		TableName: &tableName,
		AttributeDefinitions: []types.AttributeDefinition{
			{AttributeName: aws.String("ID"), AttributeType: types.ScalarAttributeTypeS},
		},
		KeySchema: []types.KeySchemaElement{
			{AttributeName: aws.String("ID"), KeyType: types.KeyTypeHash},
		},
		ProvisionedThroughput: &types.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(1),
			WriteCapacityUnits: aws.Int64(1),
		},
	})

	if err != nil {
		return fmt.Errorf("error making sure table '%s' exists: %w", tableName, err)
	}

	fmt.Printf("Table '%s' created successfully\n", tableName)
	return nil
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
