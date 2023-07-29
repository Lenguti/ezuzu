package paymentdb

import (
	"context"
	"fmt"

	"github.com/lenguti/ezuzu/business/core/payment"
	"github.com/lenguti/ezuzu/business/data/db"
)

// Store - manages the set of apis for payment database access.
type Store struct {
	db *db.DB
}

// NewStore - constructs the api for data access.
func NewStore(db *db.DB) *Store {
	return &Store{
		db: db,
	}
}

// Create - will insert a new payment record.
func (s *Store) Create(ctx context.Context, p payment.Payment) error {
	dbPayment := toDBPayment(p)
	const q = `
	INSERT INTO payments (
		id,
		tenant_id,
		invoice_id,
		amount,
		created_at,
		updated_at
	) VALUES (
		:id,
		:tenant_id,
		:invoice_id,
		:amount,
		:created_at,
		:updated_at
	)
	`
	if err := s.db.Exec(ctx, q, dbPayment); err != nil {
		return fmt.Errorf("create: failed to create payment: %w", err)
	}
	return nil
}

// ListByInvoice - will query all payments for provided invoice id.
func (s *Store) ListByInvoice(ctx context.Context, invoiceID string) ([]payment.Payment, error) {
	const q = `
	SELECT *
	FROM payments
	WHERE invoice_id = $1
	`
	var out []dbPayment
	if err := s.db.List(ctx, &out, q, invoiceID); err != nil {
		return nil, fmt.Errorf("list: failed to query payments: %w", err)
	}
	return toCorePayments(out), nil
}
