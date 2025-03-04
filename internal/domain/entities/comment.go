package entities

import (
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type Comment struct {
	ID          string `json:"id" dynamodbav:"id"`
	ComponentID string `json:"componentId" dynamodbav:"componentId"`
	UserID      string `json:"userId" dynamodbav:"userId"`
	Message     string `json:"message" dynamodbav:"message"`
	UpdatedDate string `json:"updatedDate" dynamodbav:"updatedDate"`
}

func (u *Comment) Marshal() (map[string]types.AttributeValue, error) {
	return attributevalue.MarshalMap(u)
}

func (u *Comment) Unmarshal(m map[string]types.AttributeValue) error {
	return attributevalue.UnmarshalMap(m, u)
}
