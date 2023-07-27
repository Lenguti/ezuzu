package tennantdb

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/lenguti/ezuzu/business/core"
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

// Get - will fetch a tennant by its id.
func (s *Store) Get(ctx context.Context, id string) (tennant.Tennant, error) {
	const q = `
	SELECT *
	FROM tennant
	WHERE id = $1
	`
	var out dbTennant
	if err := s.db.Get(ctx, &out, q, id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return tennant.Tennant{}, core.ErrNotFound
		}
		return tennant.Tennant{}, fmt.Errorf("get: failed to fetch tennant: %w", err)
	}
	return toCoreTennant(out), nil
}

// List - will list all tennants for the provided property id.
func (s *Store) List(ctx context.Context, propertyID string) ([]tennant.Tennant, error) {
	const q = `
	SELECT *
	FROM tennant
	WHERE property_id = $1
	`
	var out []dbTennant
	if err := s.db.List(ctx, &out, q, propertyID); err != nil {
		return nil, fmt.Errorf("list: failed to list tennants: %w", err)
	}
	return toCoreTennants(out), nil
}
