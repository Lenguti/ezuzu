package paymenthistory

import (
	"context"

	"github.com/rs/zerolog"
)

// Storer - represents the data layer behavior for payment history.
type Storer interface {
	List(ctx context.Context, tenantID string) ([]PaymentHistory, error)
}

// Core - represents the core business logic for payment history.
type Core struct {
	store Storer
	log   zerolog.Logger
}

// NewCore - returns a new payment history core with all its components initialized.
func NewCore(store Storer, log zerolog.Logger) *Core {
	return &Core{
		store: store,
		log:   log,
	}
}
