package auth

import (
	"context"
	"fmt"
	"os"
	"time"

	"xrouting/internal/domain/entities"
	"xrouting/internal/domain/repositories"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func ComparePassword(hashedPassword, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)) == nil
}

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

	hashedPassword, err := HashPassword(user.Password)
	if err != nil {
		return fmt.Errorf("error hashing password: %v", err)
	}
	user.Password = hashedPassword

	return r.userRepo.CreateUser(ctx, "users", user)
}

func (r *authRepository) AuthenticateUser(ctx context.Context, username, password string) (string, error) {
	findUser, err := r.userRepo.ListUsers(ctx, "users")
	if err != nil {
		return "", fmt.Errorf("user not found")
	}

	var user *entities.User
	for _, u := range findUser {
		if u.Username == username {
			user = u
			break
		}
	}

	if user == nil {
		return "", fmt.Errorf("user not found")
	}

	if !ComparePassword(user.Password, password) {
		return "", fmt.Errorf("invalid credentials")
	}

	claims := JwtClaims{
		Username: username,
		Role:     user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			// ExpiresAt, never expires
			ExpiresAt: nil,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	return tokenString, err
}
