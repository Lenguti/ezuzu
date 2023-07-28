package property

import (
	"context"

	"github.com/rs/zerolog"
)

// Storer - represents the data layer behavior for property.
type Storer interface {
	Create(ctx context.Context, p Property) error
	UpdateName(ctx context.Context, p Property) error
	Get(ctx context.Context, id string) (Property, error)
	List(ctx context.Context, managerID string) ([]Property, error)
}

// Core - represents the core business logic for property.
type Core struct {
	store Storer
	log   zerolog.Logger
}

// NewCore - returns a new property core with all its components initialized.
func NewCore(store Storer, log zerolog.Logger) *Core {
	return &Core{
		store: store,
		log:   log,
	}
}
