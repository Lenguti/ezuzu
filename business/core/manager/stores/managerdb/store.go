package managerdb

import (
	"context"
	"fmt"

	"github.com/lenguti/ezuzu/business/core/manager"
	"github.com/lenguti/ezuzu/business/data/db"
)

// Store - manages the set of apis for property manager database access.
type Store struct {
	db *db.DB
}

// NewStore - constructs the api for data access.
func NewStore(db *db.DB) *Store {
	return &Store{
		db: db,
	}
}

// Create - will insert a new property manager record.
func (s *Store) Create(ctx context.Context, m manager.Manager) error {
	dbManager := toDBManager(m)
	const q = `
	INSERT INTO property_manager (
		id,
		entity,
		created_at,
		updated_at
	) VALUES (
		:id,
		:entity,
		:created_at,
		:updated_at
	)
	`
	if err := s.db.Exec(ctx, q, dbManager); err != nil {
		return fmt.Errorf("create: failed to create manager: %w", err)
	}
	return nil
}
