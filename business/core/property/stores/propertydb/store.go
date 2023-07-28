package propertydb

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/lenguti/ezuzu/business/core"
	"github.com/lenguti/ezuzu/business/core/property"
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
func (s *Store) Create(ctx context.Context, p property.Property) error {
	dbProperty := toDBProperty(p)
	const q = `
	INSERT INTO property (
		id,
		manager_id,
		address,
		name,
		type,
		unit_number,
		created_at,
		updated_at
	) VALUES (
		:id,
		:manager_id,
		:address,
		:name,
		:type,
		:unit_number,
		:created_at,
		:updated_at
	)
	`
	if err := s.db.Exec(ctx, q, dbProperty); err != nil {
		return fmt.Errorf("create: failed to create property: %w", err)
	}
	return nil
}

// UpdateName - will update a property record with a new name.
func (s *Store) UpdateName(ctx context.Context, p property.Property) error {
	dbProperty := toDBProperty(p)
	const q = `
	UPDATE property
	SET
	name = :name,
	updated_at = :updated_at
	WHERE id = :id
	`
	if err := s.db.Exec(ctx, q, dbProperty); err != nil {
		return fmt.Errorf("update name: failed to update property: %w", err)
	}
	return nil
}

// Get - will fetch a property by its id.
func (s *Store) Get(ctx context.Context, id string) (property.Property, error) {
	const q = `
	SELECT *
	FROM property
	WHERE id = $1
	`
	var out dbProperty
	if err := s.db.Get(ctx, &out, q, id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return property.Property{}, core.ErrNotFound
		}
		return property.Property{}, fmt.Errorf("get: failed to fetch property: %w", err)
	}
	return toCoreProperty(out), nil
}

// List - will list all properties for the provided manager id.
func (s *Store) List(ctx context.Context, managerID string) ([]property.Property, error) {
	const q = `
	SELECT *
	FROM property
	WHERE manager_id = $1
	`
	var out []dbProperty
	if err := s.db.List(ctx, &out, q, managerID); err != nil {
		return nil, fmt.Errorf("list: failed to list properties: %w", err)
	}
	return toCoreProperties(out), nil
}
