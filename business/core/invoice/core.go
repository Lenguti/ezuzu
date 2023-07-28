package invoice

import (
	"context"

	"github.com/rs/zerolog"
)

// Storer - represents the data layer behavior for invoices.
type Storer interface {
	Create(ctx context.Context, i Invoice) error
	Get(ctx context.Context, id string) (Invoice, error)
}

// Core - represents the core business logic for invoices.
type Core struct {
	store Storer
	log   zerolog.Logger
}

// NewCore - returns a new invoices core with all its components initialized.
func NewCore(store Storer, log zerolog.Logger) *Core {
	return &Core{
		store: store,
		log:   log,
	}
}
