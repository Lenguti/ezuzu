package tenant

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
)

// Create - will create a new tenant.
func (c *Core) Create(ctx context.Context, nt NewTenant, propertyID uuid.UUID) (Tenant, error) {
	now := time.Now().UTC()
	t := Tenant{
		ID:          uuid.New(),
		PropertyID:  propertyID,
		FirstName:   nt.FirstName,
		LastName:    nt.LastName,
		DateOfBirth: nt.DateOfBirth,
		SSN:         nt.SSN,
		Type:        nt.Type,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
	if err := c.store.Create(ctx, t); err != nil {
		return Tenant{}, fmt.Errorf("create: unable to create tenant: %w", err)
	}
	return t, nil
}

// Update - will update a tenant.
func (c *Core) Update(ctx context.Context, id uuid.UUID, ut UpdateTenant) (Tenant, error) {
	t, err := c.Get(ctx, id)
	if err != nil {
		return Tenant{}, fmt.Errorf("update: unable to fetch tenant: %w", err)
	}

	if ut.PropertyID != nil {
		p, err := c.pc.Get(ctx, *ut.PropertyID)
		if err != nil {
			return Tenant{}, fmt.Errorf("update: unable to fetch property: %w", err)
		}
		t.PropertyID = p.ID
	}

	if ut.Type != nil {
		t.Type = *ut.Type
	}
	t.UpdatedAt = time.Now().UTC()
	if err := c.store.Update(ctx, t); err != nil {
		return Tenant{}, fmt.Errorf("update: unable to update tenant: %w", err)
	}
	return t, nil
}

// Get - will get a tenant by the provided id.
func (c *Core) Get(ctx context.Context, id uuid.UUID) (Tenant, error) {
	t, err := c.store.Get(ctx, id.String())
	if err != nil {
		return Tenant{}, fmt.Errorf("get: unable to get tenant: %w", err)
	}
	return t, nil
}

// List - will list tenants for the provided property id.
func (c *Core) List(ctx context.Context, propertyID uuid.UUID) ([]Tenant, error) {
	tts, err := c.store.List(ctx, propertyID.String())
	if err != nil {
		return nil, fmt.Errorf("list: unable to list tenants: %w", err)
	}
	return tts, nil
}
