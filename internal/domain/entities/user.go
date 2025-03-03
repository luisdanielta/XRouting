package entities

import (
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

var (
	ErrDuplicateEmail    = errors.New("a user with that email already exists")
	ErrDuplicateUsername = errors.New("a user with that username already exists")
)

type User struct {
	ID       int      `json:"id"`
	Username string   `json:"username"`
	Email    string   `json:"email"`
	Password Password `json:"-"`
	IsActive bool     `json:"is_active"`
	Role     string   `json:"role"`
}

func (u User) GetID() string {
	return fmt.Sprintf("%d", u.ID)
}

type Password struct {
	text *string
	hash []byte
}

func (p *Password) Set(text string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(text), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	p.text = &text
	p.hash = hash

	return nil
}

func (p *Password) Compare(text string) error {
	return bcrypt.CompareHashAndPassword(p.hash, []byte(text))
}
