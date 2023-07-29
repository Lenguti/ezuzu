package paymenthistory

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

// List - will fetch payment history for the given tenant id.
func (c *Core) List(ctx context.Context, tenantID uuid.UUID) ([]PaymentHistory, error) {
	phs, err := c.store.List(ctx, tenantID.String())
	if err != nil {
		return nil, fmt.Errorf("list: unable to list payment history: %w", err)
	}
	return phs, nil
}
