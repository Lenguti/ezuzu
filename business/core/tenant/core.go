package tenant

import (
	"context"

	"github.com/lenguti/ezuzu/business/core/property"
	"github.com/rs/zerolog"
)

// Storer - represents the data layer behavior for tenant.
type Storer interface {
	Create(ctx context.Context, t Tenant) error
	Update(ctx context.Context, t Tenant) error
	Get(ctx context.Context, id string) (Tenant, error)
	List(ctx context.Context, propertyID string) ([]Tenant, error)
}

// Core - represents the core business logic for tenant.
type Core struct {
	pc    *property.Core
	store Storer
	log   zerolog.Logger
}

// NewCore - returns a new tenant core with all its components initialized.
func NewCore(store Storer, pc *property.Core, log zerolog.Logger) *Core {
	return &Core{
		pc:    pc,
		store: store,
		log:   log,
	}
}
