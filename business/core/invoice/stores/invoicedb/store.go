package invoicedb

import (
	"context"
	"fmt"

	"github.com/lenguti/ezuzu/business/core/invoice"
	"github.com/lenguti/ezuzu/business/data/db"
)

// Store - manages the set of apis for property manager database access.
type Store struct {
	db *db.DB
}

// NewStore - constructs the api for data access.
func NewStore(db *db.DB) *Store {
	return &Store{
		db: db,
	}
}

// Create - will insert a new invoice record.
func (s *Store) Create(ctx context.Context, i invoice.Invoice) error {
	dbInvoice := toDBInvoice(i)
	const q = `
	INSERT INTO invoices (
		id,
		manager_id,
		property_id,
		tenant_id,
		amount,
		due_date,
		created_at,
		updated_at
	) VALUES (
		:id,
		:manager_id,
		:property_id,
		:tenant_id,
		:amount,
		:due_date,
		:created_at,
		:updated_at
	)
	`
	if err := s.db.Exec(ctx, q, dbInvoice); err != nil {
		return fmt.Errorf("create: failed to create invoice: %w", err)
	}
	return nil
}
