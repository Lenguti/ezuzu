package invoice

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
)

// Create - will create a new invoice.
func (c *Core) Create(ctx context.Context, ni NewInvoice) (Invoice, error) {
	now := time.Now().UTC()
	i := Invoice{
		ID:         uuid.New(),
		ManagerID:  ni.ManagerID,
		PropertyID: ni.PropertyID,
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

// Get - will fetch an invoice by the provided id.
func (c *Core) Get(ctx context.Context, id uuid.UUID) (Invoice, error) {
	i, err := c.store.Get(ctx, id.String())
	if err != nil {
		return Invoice{}, fmt.Errorf("get: unable to get invoice: %w", err)
	}
	return i, nil
}
