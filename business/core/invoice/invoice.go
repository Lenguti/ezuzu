package invoice

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
)

// Create - will create a new invoice.
func (c *Core) Create(ctx context.Context, ni NewInvoice, managerID, propertyID uuid.UUID) (Invoice, error) {
	now := time.Now().UTC()
	i := Invoice{
		ID:         uuid.New(),
		ManagerID:  managerID,
		PropertyID: propertyID,
		TenantID:   ni.TenantID,
		Amount:     ni.Amount,
		DueDate:    ni.DueDate,
		CreatedAt:  now,
		UpdatedAt:  now,
	}
	if err := c.store.Create(ctx, i); err != nil {
		return Invoice{}, fmt.Errorf("create: unable to create invoice: %w", err)
	}
	return i, nil
}
