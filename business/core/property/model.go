package property

import (
	"time"

	"github.com/google/uuid"
)

// Property - represents a business domain property.
type Property struct {
	ID         uuid.UUID
	ManagerID  uuid.UUID
	Address    string
	Name       string
	Rent       float64
	Type       Type
	UnitNumber *int
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

// NewProperty - represents fields needed to create a new property.
type NewProperty struct {
	Address    string
	Name       string
	Rent       float64
	Type       Type
	UnitNumber *int
}

// UpdateProperty - represents fields needed to update an existing property.
type UpdateProperty struct {
	Name *string
	Rent *float64
}
