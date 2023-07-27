package manager

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
)

// Create - will create a new property manager.
func (c *Core) Create(ctx context.Context, nm NewManager) (Manager, error) {
	now := time.Now().UTC()
	m := Manager{
		ID:        uuid.New(),
		Entity:    nm.Entity,
		CreatedAt: now,
		UpdatedAt: now,
	}
	if err := c.store.Create(ctx, m); err != nil {
		return Manager{}, fmt.Errorf("create: unable to create property manager: %w", err)
	}
	return m, nil
}
