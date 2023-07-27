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
