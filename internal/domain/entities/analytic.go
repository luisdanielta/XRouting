package entities

import (
	"time"
)

// Analytic represents an analytic record
type Analytic struct {
	ID         string               `json:"id" dynamodbav:"id"`
	MetricID   string               `json:"metricId" dynamodbav:"metricId"`
	MetricName string               `json:"metricName" dynamodbav:"metricName"`
	Data       []map[string]float64 `json:"data" dynamodbav:"data"`
	UpdatedAt  time.Time            `json:"updatedAt" dynamodbav:"updatedAt"`
}
