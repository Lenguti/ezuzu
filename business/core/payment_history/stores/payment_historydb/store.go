package paymenthistorydb

import (
	"context"
	"fmt"

	paymenthistory "github.com/lenguti/ezuzu/business/core/payment_history"
	"github.com/lenguti/ezuzu/business/data/db"
)

// Store - manages the set of apis for payment history database access.
type Store struct {
	db *db.DB
}

// NewStore - constructs the api for data access.
func NewStore(db *db.DB) *Store {
	return &Store{
		db: db,
	}
}

// List - will query payment history for the provided tenant.
func (s *Store) List(ctx context.Context, tenantID string) ([]paymenthistory.PaymentHistory, error) {
	const q = `
	SELECT *
	FROM payment_history
	WHERE id = $1
	ORDER BY earliest_payment DESC
	`
	var out []dbPaymentHistory
	if err := s.db.List(ctx, &out, q, tenantID); err != nil {
		return nil, fmt.Errorf("list: failed to query payment history: %w", err)
	}
	return toCorePaymentHistories(out), nil
}
