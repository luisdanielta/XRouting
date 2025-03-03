package entities

type ComponentType string

type ComponentStatus string

type ComponentCategory string

const (
	Core      ComponentType = "core"
	Capsule   ComponentType = "capsule"
	Engine    ComponentType = "engine"
	Ship      ComponentType = "ship"
	Launchpad ComponentType = "launchpad"
	Rocket    ComponentType = "rocket"
)

const (
	Active    ComponentStatus = "active"
	Inactive  ComponentStatus = "inactive"
	Damaged   ComponentStatus = "damaged"
	Destroyed ComponentStatus = "destroyed"
)

const (
	Fixed    ComponentCategory = "fixed"
	Launched ComponentCategory = "launched"
)

type Component struct {
	ID               string            `dynamodbav:"id" json:"id"`
	ComponentID      string            `dynamodbav:"component_id" json:"component_id"`
	Name             string            `dynamodbav:"name" json:"name"`
	Type             ComponentType     `dynamodbav:"type" json:"type"`
	Status           ComponentStatus   `dynamodbav:"status" json:"status"`
	Category         ComponentCategory `dynamodbav:"category" json:"category"`
	Missions         []string          `dynamodbav:"missions" json:"missions"`
	HasSubcomponents bool
}
