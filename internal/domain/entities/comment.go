package entities

type Comment struct {
	ID          string `json:"id" dynamodbav:"id"`
	ComponentID string `json:"componentId" dynamodbav:"componentId"`
	UserID      string `json:"userId" dynamodbav:"userId"`
	Message     string `json:"message" dynamodbav:"message"`
	UpdatedDate string `json:"updatedDate" dynamodbav:"updatedDate"`
}
