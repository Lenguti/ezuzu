package payment

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/lenguti/ezuzu/business/core"
)

// Create - will create a new payment.
func (c *Core) Create(ctx context.Context, np NewPayment) (Payment, error) {
	i, err := c.ic.Get(ctx, np.InvoiceID)
	if err != nil {
		return Payment{}, fmt.Errorf("create: unable to fetch invoice: %w", err)
	}

	if time.Now().UTC().After(i.DueDate) {
		return Payment{}, core.ErrPastDuePayment
	}

	ps, err := c.ListByInvoice(ctx, i.ID)
	if err != nil {
		return Payment{}, fmt.Errorf("create: unable to list payments for invoice: %w", err)
	}

	var currentPaid float64
	for _, p := range ps {
		currentPaid += p.Amount
	}

	payDiff := i.Amount - currentPaid
	if np.Amount > payDiff {
		return Payment{}, core.ErrPaymentConflict
	}

	now := time.Now().UTC()
	p := Payment{
		ID:        uuid.New(),
		TenantID:  np.TenantID,
		InvoiceID: np.InvoiceID,
		Amount:    np.Amount,
		CreatedAt: now,
		UpdatedAt: now,
	}
	if err := c.store.Create(ctx, p); err != nil {
		return Payment{}, fmt.Errorf("create: unable to create payment: %w", err)
	}
	return p, nil
}

// List - will fetch all payments for a provided invoice.
func (c *Core) ListByInvoice(ctx context.Context, invoiceID uuid.UUID) ([]Payment, error) {
	ps, err := c.store.ListByInvoice(ctx, invoiceID.String())
	if err != nil {
		return nil, fmt.Errorf("create: unable to list payments: %w", err)
	}
	return ps, nil
}
