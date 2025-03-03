package entities

import "time"

// FailureRiskLevel defines levels of failure risks
type FailureRiskLevel string

const (
	Low    FailureRiskLevel = "low"
	Medium FailureRiskLevel = "medium"
	High   FailureRiskLevel = "high"
)

// FailurePrediction represents a prediction of failure
type FailurePrediction struct {
	RiskLevel            FailureRiskLevel `json:"riskLevel" dynamodbav:"riskLevel"`
	EstimatedFailureDate *time.Time       `json:"estimatedFailureDate,omitempty" dynamodbav:"estimatedFailureDate,omitempty"`
	Reason               string           `json:"reason" dynamodbav:"reason"`
}

// Maintenance represents maintenance data
type Maintenance struct {
	ID                    string             `json:"id" dynamodbav:"id"`
	MaintenanceID         string             `json:"maintenanceId" dynamodbav:"maintenanceId"`
	ComponentID           string             `json:"componentId" dynamodbav:"componentId"`
	SubcomponentsAffected []string           `json:"subcomponentsAffected" dynamodbav:"subcomponentsAffected"`
	Date                  time.Time          `json:"date" dynamodbav:"date"`
	PerformedBy           string             `json:"performedBy" dynamodbav:"performedBy"`
	Cost                  float64            `json:"cost" dynamodbav:"cost"`
	Description           *string            `json:"description,omitempty" dynamodbav:"description,omitempty"`
	NextMaintenance       *time.Time         `json:"nextMaintenance,omitempty" dynamodbav:"nextMaintenance,omitempty"`
	FailurePrediction     *FailurePrediction `json:"failurePrediction,omitempty" dynamodbav:"failurePrediction,omitempty"`
}
