package invoicedb

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/lenguti/ezuzu/business/core"
	"github.com/lenguti/ezuzu/business/core/invoice"
	"github.com/lenguti/ezuzu/business/data/db"
)

// Store - manages the set of apis for invoice manager database access.
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

// Get - will fetch an invoice by its id.
func (s *Store) Get(ctx context.Context, id string) (invoice.Invoice, error) {
	const q = `
	SELECT *
	FROM invoices
	WHERE id = $1
	`
	var out dbInvoice
	if err := s.db.Get(ctx, &out, q, id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return invoice.Invoice{}, core.ErrNotFound
		}
		return invoice.Invoice{}, fmt.Errorf("get: failed to fetch invoice: %w", err)
	}
	return toCoreInvoice(out), nil
}
