package invoice

import (
	"time"

	"github.com/google/uuid"
)

// Invoice - represents a business domain invoice.
type Invoice struct {
	ID         uuid.UUID
	ManagerID  uuid.UUID
	PropertyID uuid.UUID
	TenantID   uuid.UUID
	Amount     float64
	DueDate    time.Time
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

// NewInvoice - represents fields needed to create a new invoice.
type NewInvoice struct {
	TenantID uuid.UUID
	Amount   float64
	DueDate  time.Time
}
