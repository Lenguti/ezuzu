package payment

import (
	"time"

	"github.com/google/uuid"
)

// Payment - represents a business domain payment.
type Payment struct {
	ID        uuid.UUID
	TenantID  uuid.UUID
	InvoiceID uuid.UUID
	Amount    float64
	CreatedAt time.Time
	UpdatedAt time.Time
}

// NewPayment - represents fields needed to create a new payment.
type NewPayment struct {
	TenantID  uuid.UUID
	InvoiceID uuid.UUID
	Amount    float64
}
