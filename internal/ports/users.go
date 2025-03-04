package ports

import (
	"context"
	"xrouting/internal/domain/entities"
	"xrouting/internal/domain/repositories"
)

type UserService interface {
	RegisterUser(ctx context.Context, tableName string, user *entities.User) error
	FindUser(ctx context.Context, tableName string, userID string) (*entities.User, error)
	RemoveUser(ctx context.Context, tableName string, userID string) error
}

type userService struct {
	userRepo repositories.UserRepository
}

func NewUserService(userRepo repositories.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

func (s *userService) RegisterUser(ctx context.Context, tableName string, user *entities.User) error {
	return s.userRepo.CreateUser(ctx, tableName, user)
}

func (s *userService) FindUser(ctx context.Context, tableName string, userID string) (*entities.User, error) {
	return s.userRepo.GetUser(ctx, tableName, userID)
}

func (s *userService) RemoveUser(ctx context.Context, tableName string, userID string) error {
	return s.userRepo.DeleteUser(ctx, tableName, userID)
}
