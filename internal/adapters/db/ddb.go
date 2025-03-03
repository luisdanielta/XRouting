package db

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

// Marshaler defines how to convert an entity to/from a DynamoDB attribute map.
type Marshaler interface {
	Marshal() (map[string]types.AttributeValue, error)
	Unmarshal(map[string]types.AttributeValue) error
}

// DynamoDBRepository defines basic CRUD operations for objects
// that implement Marshaler.
type DynamoDBRepository interface {
	PutItem(ctx context.Context, tableName string, item Marshaler) error
	GetItem(ctx context.Context, tableName string, key map[string]types.AttributeValue, out Marshaler) error
	DeleteItem(ctx context.Context, tableName string, key map[string]types.AttributeValue) error
	UpdateItem(
		ctx context.Context,
		tableName string,
		key map[string]types.AttributeValue,
		updateExpression string,
		expressionValues map[string]types.AttributeValue,
	) error
}

// DynamoDBClient implements the DynamoDBRepository interface.
// It requires an AWS config to create the service client.
type DynamoDBClient struct {
	client *dynamodb.Client
}

// NewDynamoDBClient creates a new DynamoDB client.
func NewDynamoDBClient(cfg aws.Config) *DynamoDBClient {
	return &DynamoDBClient{
		client: dynamodb.NewFromConfig(cfg),
	}
}

func (c *DynamoDBClient) PutItem(ctx context.Context, tableName string, item Marshaler) error {
	attrMap, err := item.Marshal()
	if err != nil {
		return err
	}
	_, err = c.client.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: &tableName,
		Item:      attrMap,
	})
	return err
}

func (c *DynamoDBClient) GetItem(
	ctx context.Context,
	tableName string,
	key map[string]types.AttributeValue,
	out Marshaler,
) error {
	resp, err := c.client.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: &tableName,
		Key:       key,
	})
	if err != nil {
		return err
	}
	return out.Unmarshal(resp.Item)
}

func (c *DynamoDBClient) DeleteItem(
	ctx context.Context,
	tableName string,
	key map[string]types.AttributeValue,
) error {
	_, err := c.client.DeleteItem(ctx, &dynamodb.DeleteItemInput{
		TableName: &tableName,
		Key:       key,
	})
	return err
}

func (c *DynamoDBClient) UpdateItem(
	ctx context.Context,
	tableName string,
	key map[string]types.AttributeValue,
	updateExpression string,
	expressionValues map[string]types.AttributeValue,
) error {
	_, err := c.client.UpdateItem(ctx, &dynamodb.UpdateItemInput{
		TableName:                 &tableName,
		Key:                       key,
		UpdateExpression:          &updateExpression,
		ExpressionAttributeValues: expressionValues,
	})
	return err
}
