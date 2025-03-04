package entities

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

// Analytic represents an analytic record
type Analytic struct {
	ID         string           `json:"id" dynamodbav:"id"`
	MetricID   string           `json:"metricId" dynamodbav:"metricId"`
	MetricName string           `json:"metricName" dynamodbav:"metricName"`
	Data       []map[string]any `json:"data" dynamodbav:"data"`
	UpdatedAt  time.Time        `json:"updatedAt" dynamodbav:"updatedAt"`
}

func (u *Analytic) Marshal() (map[string]types.AttributeValue, error) {
	return attributevalue.MarshalMap(u)
}

func (u *Analytic) Unmarshal(m map[string]types.AttributeValue) error {
	return attributevalue.UnmarshalMap(m, u)
}
