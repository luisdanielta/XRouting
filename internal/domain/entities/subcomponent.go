package entities

import "time"

// SubComponentStatus defines the status of a subcomponent
type SubComponentStatus string

const (
	Good       SubComponentStatus = "good"
	NeedReview SubComponentStatus = "needReview"
)

// SubComponent represents a subcomponent entity
type SubComponent struct {
	ID                string             `json:"id" dynamodbav:"id"`
	SubComponentID    string             `json:"subComponentId" dynamodbav:"subComponentId"`
	ParentComponentID string             `json:"parentComponentId" dynamodbav:"parentComponentId"`
	Name              string             `json:"name" dynamodbav:"name"`
	Status            SubComponentStatus `json:"status" dynamodbav:"status"`
}

// FailureType defines different types of failures
type FailureType string

const (
	Wear        FailureType = "desgaste"
	Overheating FailureType = "sobrecalentamiento"
	Vibrations  FailureType = "vibraciones"
)

// FailureStatus defines the status of failures
type FailureStatus string

const (
	Pending     FailureStatus = "pending"
	UnderRepair FailureStatus = "underRepair"
	Resolved    FailureStatus = "resolved"
)

// SubComponentFailure represents a failure detected in a subcomponent
type SubComponentFailure struct {
	ID             string           `json:"id" dynamodbav:"id"`
	SubComponentID string           `json:"subComponentId" dynamodbav:"subComponentId"`
	DetectedAt     time.Time        `json:"detectedAt" dynamodbav:"detectedAt"`
	FailureType    FailureType      `json:"failureType" dynamodbav:"failureType"`
	Severity       FailureRiskLevel `json:"severity" dynamodbav:"severity"`
	ReportedBy     string           `json:"reportedBy" dynamodbav:"reportedBy"`
	Status         FailureStatus    `json:"status" dynamodbav:"status"`
	MaintenanceID  *string          `json:"maintenanceId,omitempty" dynamodbav:"maintenanceId,omitempty"`
}
