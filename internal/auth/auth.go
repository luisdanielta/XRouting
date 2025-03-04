package auth

import (
	"context"
	"fmt"
	"xrouting/internal/domain/entities"
	"xrouting/internal/domain/repositories"

	"github.com/golang-jwt/jwt/v5"
)

type Authenticator interface {
	GenerateToken(claims jwt.Claims) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}

type AuthRepository interface {
	RegisterUser(ctx context.Context, user *entities.User) error
	AuthenticateUser(ctx context.Context, username, password string) (string, error)
}

type authRepository struct {
	userRepo repositories.UserRepository
}

func NewAuthRepository(userRepo repositories.UserRepository) AuthRepository {
	return &authRepository{
		userRepo: userRepo,
	}
}

func (r *authRepository) RegisterUser(ctx context.Context, user *entities.User) error {

	users, err := r.userRepo.ListUsers(ctx, "users")

	if err != nil {
		return fmt.Errorf("error listing users: %v", err)
	}

	for _, u := range users {
		if u.Username == user.Username {
			return fmt.Errorf("user %s already exists", user.Username)
		}
	}

	return r.userRepo.CreateUser(ctx, "users", user)
}

func (r *authRepository) AuthenticateUser(ctx context.Context, username, password string) (string, error) {
	user, err := r.userRepo.ListUsers(ctx, "users")
	if err != nil {
		return "", fmt.Errorf("error listing users: %v", err)
	}

	for _, u := range user {
		if u.Username == username && u.Password == password {
			return u.Username, nil
		}
	}

	return "", fmt.Errorf("invalid username or password")
}
