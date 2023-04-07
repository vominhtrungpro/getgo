package model

import "time"

// CategoryStatus represents the status of the category
type CategoryStatus string

const (
	// CategoryStatusActive means the category is active
	CategoryStatusActive CategoryStatus = "ACTIVE"
	// CategoryStatusInactive means the category is inactive
	CategoryStatusInactive CategoryStatus = "INACTIVE"
	// CategoryStatusDeleted means the category is deleted. This is for archival only
	CategoryStatusDeleted CategoryStatus = "DELETED"
)

// AllcategoryStatus returns list of all CategoryStatus
var AllcategoryStatus = []CategoryStatus{CategoryStatusActive, CategoryStatusInactive, CategoryStatusDeleted}

// String converts to string value
func (p CategoryStatus) String() string {
	return string(p)
}

// IsValid checks if category status is valid
func (p CategoryStatus) IsValid() bool {
	return p == CategoryStatusActive || p == CategoryStatusInactive || p == CategoryStatusDeleted
}

// Category represents the business model of category
type Category struct {
	ID          int64
	ExternalID  string
	Description string
	Name        string
	Status      CategoryStatus
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
