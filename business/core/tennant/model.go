package tennant

import (
	"time"

	"github.com/google/uuid"
)

// Tennant - represents a business domain tennant.
type Tennant struct {
	ID          uuid.UUID
	PropertyID  uuid.UUID
	FirstName   string
	LastName    string
	DateOfBirth string
	SSN         int
	Type        Type
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// NewTennant - represents fields needed to create a new tennant.
type NewTennant struct {
	FirstName   string
	LastName    string
	DateOfBirth string
	SSN         int
	Type        Type
}
