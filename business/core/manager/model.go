package manager

import (
	"time"

	"github.com/google/uuid"
)

// Manager - represents a business domain property manager.
type Manager struct {
	ID        uuid.UUID
	Entity    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// NewManager - represents fields needed to create a new manager.
type NewManager struct {
	Entity string
}
