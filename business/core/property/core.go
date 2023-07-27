package property

import (
	"context"

	"github.com/rs/zerolog"
)

// Storer - represents the data layer behavior for cages.
type Storer interface {
	Create(ctx context.Context, p Property) error
	Get(ctx context.Context, id string) (Property, error)
	List(ctx context.Context, managerID string) ([]Property, error)
}

// Core - represents the core business logic for cages.
type Core struct {
	store Storer
	log   zerolog.Logger
}

// NewCore - returns a new cage core with all its components initialized.
func NewCore(store Storer, log zerolog.Logger) *Core {
	return &Core{
		store: store,
		log:   log,
	}
}
