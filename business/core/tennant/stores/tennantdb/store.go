package tennantdb

import (
	"context"
	"fmt"

	"github.com/lenguti/ezuzu/business/core/tennant"
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

// Create - will insert a new property record.
func (s *Store) Create(ctx context.Context, t tennant.Tennant) error {
	dbTennant := toDBTennant(t)
	const q = `
	INSERT INTO tennant (
		id,
		property_id,
		first_name,
		last_name,
		date_of_birth,
		ssn,
		type,
		created_at,
		updated_at
	) VALUES (
		:id,
		:property_id,
		:first_name,
		:last_name,
		:date_of_birth,
		:ssn,
		:type,
		:created_at,
		:updated_at
	)
	`
	if err := s.db.Exec(ctx, q, dbTennant); err != nil {
		return fmt.Errorf("create: failed to create tennant: %w", err)
	}
	return nil
}
