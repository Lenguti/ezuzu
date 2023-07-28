package property

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
)

// Create - will create a new property.
func (c *Core) Create(ctx context.Context, np NewProperty, managerID uuid.UUID) (Property, error) {
	now := time.Now().UTC()
	p := Property{
		ID:         uuid.New(),
		ManagerID:  managerID,
		Address:    np.Address,
		Name:       np.Name,
		Type:       np.Type,
		UnitNumber: np.UnitNumber,
		CreatedAt:  now,
		UpdatedAt:  now,
	}
	if err := c.store.Create(ctx, p); err != nil {
		return Property{}, fmt.Errorf("create: unable to create property: %w", err)
	}
	return p, nil
}

// UpdateName - will update the name of a property.
func (c *Core) UpdateName(ctx context.Context, id uuid.UUID, name string) (Property, error) {
	p, err := c.Get(ctx, id)
	if err != nil {
		return Property{}, fmt.Errorf("update name: unable to fetch property: %w", err)
	}

	now := time.Now().UTC()
	p.Name = name
	p.UpdatedAt = now
	if err := c.store.UpdateName(ctx, p); err != nil {
		return Property{}, fmt.Errorf("update name: unable to update property: %w", err)
	}
	return p, nil
}

// Get - will get a property by the provided id.
func (c *Core) Get(ctx context.Context, id uuid.UUID) (Property, error) {
	p, err := c.store.Get(ctx, id.String())
	if err != nil {
		return Property{}, fmt.Errorf("get: unable to get property: %w", err)
	}
	return p, nil
}

// List - will list properties for the provided manager id.
func (c *Core) List(ctx context.Context, managerID uuid.UUID) ([]Property, error) {
	pps, err := c.store.List(ctx, managerID.String())
	if err != nil {
		return nil, fmt.Errorf("list: unable to list properties: %w", err)
	}
	return pps, nil
}
