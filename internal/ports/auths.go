package ports

import (
	"context"
	"xrouting/internal/auth"
	"xrouting/internal/domain/entities"
)

type AuthService interface {
	SignUp(ctx context.Context, user *entities.User) error
	SignIn(ctx context.Context, username, password string) (string, error)
}

type authService struct {
	authRepo auth.AuthRepository
}

func NewAuthService(authRepo auth.AuthRepository) AuthService {
	return &authService{authRepo: authRepo}
}

func (s *authService) SignUp(ctx context.Context, user *entities.User) error {
	return s.authRepo.RegisterUser(ctx, user)
}

func (s *authService) SignIn(ctx context.Context, username, password string) (string, error) {
	return s.authRepo.AuthenticateUser(ctx, username, password)
}
