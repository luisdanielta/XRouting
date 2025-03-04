package repositories

import (
	"context"
	"xrouting/internal/adapters/db"
	"xrouting/internal/domain/entities"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type CommentRepository interface {
	CreateComment(ctx context.Context, tableName string, comment *entities.Comment) error
	GetComment(ctx context.Context, tableName string, commentID string) (*entities.Comment, error)
	DeleteComment(ctx context.Context, tableName string, commentID string) error
}

type commentRepository struct {
	dynamo *db.DynamoDBClient
}

func NewCommentRepository(dynamo *db.DynamoDBClient) CommentRepository {
	return &commentRepository{dynamo: dynamo}
}

func (r *commentRepository) CreateComment(ctx context.Context, tableName string, comment *entities.Comment) error {
	return r.dynamo.PutItem(ctx, tableName, comment)
}

func (r *commentRepository) GetComment(ctx context.Context, tableName string, commentID string) (*entities.Comment, error) {
	comment := &entities.Comment{}

	err := r.dynamo.GetItem(ctx, tableName, map[string]types.AttributeValue{
		"id": &types.AttributeValueMemberS{Value: commentID},
	}, comment)

	if err != nil {
		return nil, err
	}

	return comment, nil
}

func (r *commentRepository) DeleteComment(ctx context.Context, tableName string, commentID string) error {
	return r.dynamo.DeleteItem(ctx, tableName, map[string]types.AttributeValue{
		"id": &types.AttributeValueMemberS{Value: commentID},
	})
}
