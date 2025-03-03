package repositories

import (
	"context"
	"xrouting/internal/adapters/db"
	"xrouting/internal/domain/entities"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type UserRepository interface {
	CreateUser(ctx context.Context, tableName string, user *entities.User) error
	GetUser(ctx context.Context, tableName string, userID string) (*entities.User, error)
	DeleteUser(ctx context.Context, tableName string, userID string) error
}

type userRepository struct {
	dynamo db.DynamoDBRepository
}

func NewUserRepository(dynamo db.DynamoDBRepository) UserRepository {
	return &userRepository{dynamo: dynamo}
}

func (r *userRepository) CreateUser(ctx context.Context, tableName string, user *entities.User) error {
	if err := r.dynamo.PutItem(ctx, tableName, user); err != nil {
		return err
	}

	return nil
}
func (r *userRepository) GetUser(ctx context.Context, tableName string, userID string) (*entities.User, error) {
	user := &entities.User{}

	err := r.dynamo.GetItem(ctx, tableName, map[string]types.AttributeValue{
		"id": &types.AttributeValueMemberS{Value: userID},
	}, user)

	if err != nil {
		return nil, err
	}

	return user, nil
}
func (r *userRepository) DeleteUser(ctx context.Context, tableName string, userID string) error {
	return r.dynamo.DeleteItem(ctx, tableName, map[string]types.AttributeValue{
		"id": &types.AttributeValueMemberS{Value: userID},
	})
}
