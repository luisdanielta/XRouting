package entities

import (
	"errors"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrDuplicateEmail    = errors.New("a user with that email already exists")
	ErrDuplicateUsername = errors.New("a user with that username already exists")
)

type User struct {
	ID       string `dynamodbav:"id" json:"id,omitempty"`
	Username string `dynamodbav:"username" json:"username"`
	Email    string `dynamodbav:"email" json:"email"`
	Password string `dynamodbav:"password" json:"password,omitempty"`
	IsActive bool   `dynamodbav:"is_active" json:"is_active"`
	Role     string `dynamodbav:"role" json:"role"`
}

// Marshal converts the User struct into a DynamoDB attribute map.
func (u *User) Marshal() (map[string]types.AttributeValue, error) {
	return attributevalue.MarshalMap(u)
}

func (u *User) Unmarshal(m map[string]types.AttributeValue) error {
	return attributevalue.UnmarshalMap(m, u)
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
