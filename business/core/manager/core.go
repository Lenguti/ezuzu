package manager

import (
	"context"

	"github.com/rs/zerolog"
)

// Storer - represents the data layer behavior for property manager.
type Storer interface {
	Create(ctx context.Context, m Manager) error
}

// Core - represents the core business logic for property manager.
type Core struct {
	store Storer
	log   zerolog.Logger
}

// NewCore - returns a new property manager core with all its components initialized.
func NewCore(store Storer, log zerolog.Logger) *Core {
	return &Core{
		store: store,
		log:   log,
	}
}
