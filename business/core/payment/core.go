package payment

import (
	"context"

	"github.com/lenguti/ezuzu/business/core/invoice"
	"github.com/rs/zerolog"
)

// Storer - represents the data layer behavior for payments.
type Storer interface {
	Create(ctx context.Context, p Payment) error
	ListByInvoice(ctx context.Context, invoiceID string) ([]Payment, error)
}

// Core - represents the core business logic for payments.
type Core struct {
	store Storer
	ic    *invoice.Core
	log   zerolog.Logger
}

// NewCore - returns a new payments core with all its components initialized.
func NewCore(store Storer, ic *invoice.Core, log zerolog.Logger) *Core {
	return &Core{
		ic:    ic,
		store: store,
		log:   log,
	}
}
