package ports

import (
	"context"
	"fmt"
	"xrouting/internal/adapters/db"
	"xrouting/internal/domain/entities"
)

type UserService struct {
	itemService *db.ItemService
	mapper      db.DynamoDBMapper[entities.User]
}

func NewUserService(itemService *db.ItemService) *UserService {
	return &UserService{
		itemService: itemService,
		mapper:      db.DynamoDBMapper[entities.User]{},
	}
}

func (s *UserService) CreateUser(ctx context.Context, user entities.User) error {
	item, err := s.mapper.ToDynamoDBMap(user)
	if err != nil {
		return fmt.Errorf("error converting user to DynamoDB: %w", err)
	}
	return s.itemService.CreateItem(ctx, "users", item)
}

func (s *UserService) GetUser(ctx context.Context, id string) (*entities.User, error) {
	key := s.mapper.ToDynamoDBKeyID(id)
	item, err := s.itemService.ReadItem(ctx, "users", key)
	if err != nil {
		return nil, fmt.Errorf("error getting user: %w", err)
	}
	user, err := s.mapper.FromDynamoDBMap(item)
	if err != nil {
		return nil, fmt.Errorf("error converting user from DynamoDB: %w", err)
	}
	return &user, nil
}
