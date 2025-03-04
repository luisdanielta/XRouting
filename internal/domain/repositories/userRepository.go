package repositories

import (
	"context"
	"fmt"
	"xrouting/internal/adapters/db"
	"xrouting/internal/domain/entities"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type UserRepository interface {
	CreateUser(ctx context.Context, tableName string, user *entities.User) error
	GetUser(ctx context.Context, tableName string, userID string) (*entities.User, error)
	DeleteUser(ctx context.Context, tableName string, userID string) error
	ListUsers(ctx context.Context, tableName string) ([]*entities.User, error)
}

type userRepository struct {
	dynamo *db.DynamoDBClient
}

// Constructor con inyección de dependencias
func NewUserRepository(dynamo *db.DynamoDBClient) UserRepository {
	return &userRepository{dynamo: dynamo}
}

func (r *userRepository) CreateUser(ctx context.Context, tableName string, user *entities.User) error {
	return r.dynamo.PutItem(ctx, tableName, user)
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

func (r *userRepository) ListUsers(ctx context.Context, tableName string) ([]*entities.User, error) {
	users := make([]*entities.User, 0)
	items, err := r.dynamo.ScanTable(ctx, tableName)
	if err != nil {
		return nil, err
	}

	for _, item := range items {
		user := &entities.User{}
		err = user.Unmarshal(item)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal user: %w", err)
		}
		users = append(users, user)
	}

	return users, nil
}
