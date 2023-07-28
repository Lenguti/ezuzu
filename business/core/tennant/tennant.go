package tennant

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
)

// Create - will create a new tennant.
func (c *Core) Create(ctx context.Context, nt NewTennant, propertyID uuid.UUID) (Tennant, error) {
	now := time.Now().UTC()
	t := Tennant{
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
		return Tennant{}, fmt.Errorf("create: unable to create tennant: %w", err)
	}
	return t, nil
}

// Update - will update a tennant.
func (c *Core) Update(ctx context.Context, id uuid.UUID, ut UpdateTennant) (Tennant, error) {
	t, err := c.Get(ctx, id)
	if err != nil {
		return Tennant{}, fmt.Errorf("update: unable to fetch tennant: %w", err)
	}

	if ut.PropertyID != nil {
		p, err := c.pc.Get(ctx, *ut.PropertyID)
		if err != nil {
			return Tennant{}, fmt.Errorf("update: unable to fetch property: %w", err)
		}
		t.PropertyID = p.ID
	}

	if ut.Type != nil {
		t.Type = *ut.Type
	}
	t.UpdatedAt = time.Now().UTC()
	if err := c.store.Update(ctx, t); err != nil {
		return Tennant{}, fmt.Errorf("update: unable to update tennant: %w", err)
	}
	return t, nil
}

// Get - will get a tennant by the provided id.
func (c *Core) Get(ctx context.Context, id uuid.UUID) (Tennant, error) {
	t, err := c.store.Get(ctx, id.String())
	if err != nil {
		return Tennant{}, fmt.Errorf("get: unable to get tennant: %w", err)
	}
	return t, nil
}

// List - will list tennants for the provided property id.
func (c *Core) List(ctx context.Context, propertyID uuid.UUID) ([]Tennant, error) {
	tts, err := c.store.List(ctx, propertyID.String())
	if err != nil {
		return nil, fmt.Errorf("list: unable to list tennants: %w", err)
	}
	return tts, nil
}
