package tennant

import (
	"context"

	"github.com/lenguti/ezuzu/business/core/property"
	"github.com/rs/zerolog"
)

// Storer - represents the data layer behavior for tennant.
type Storer interface {
	Create(ctx context.Context, t Tennant) error
	Update(ctx context.Context, t Tennant) error
	Get(ctx context.Context, id string) (Tennant, error)
	List(ctx context.Context, propertyID string) ([]Tennant, error)
}

// Core - represents the core business logic for tennant.
type Core struct {
	pc    *property.Core
	store Storer
	log   zerolog.Logger
}

// NewCore - returns a new tennant core with all its components initialized.
func NewCore(store Storer, pc *property.Core, log zerolog.Logger) *Core {
	return &Core{
		pc:    pc,
		store: store,
		log:   log,
	}
}
