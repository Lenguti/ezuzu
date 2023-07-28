package tenant

import (
	"time"

	"github.com/google/uuid"
)

// Tenant - represents a business domain tenant.
type Tenant struct {
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

// NewTenant - represents fields needed to create a new tenant.
type NewTenant struct {
	FirstName   string
	LastName    string
	DateOfBirth string
	SSN         int
	Type        Type
}

// UpdateTenant - represents fields needed to update an existing tenant.
type UpdateTenant struct {
	PropertyID *uuid.UUID
	Type       *Type
}
