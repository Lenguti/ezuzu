package tenantdb

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/lenguti/ezuzu/business/core"
	"github.com/lenguti/ezuzu/business/core/tenant"
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
func (s *Store) Create(ctx context.Context, t tenant.Tenant) error {
	dbTenant := toDBTenant(t)
	const q = `
	INSERT INTO tenant (
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
	if err := s.db.Exec(ctx, q, dbTenant); err != nil {
		return fmt.Errorf("create: failed to create tenant: %w", err)
	}
	return nil
}

// Update - will update a tenant record with new values.
func (s *Store) Update(ctx context.Context, t tenant.Tenant) error {
	dbTenant := toDBTenant(t)
	const q = `
	UPDATE tenant
	SET
	property_id = :property_id,
	type = :type,
	updated_at = :updated_at
	WHERE id = :id
	`
	if err := s.db.Exec(ctx, q, dbTenant); err != nil {
		return fmt.Errorf("update: failed to update tenant: %w", err)
	}
	return nil
}

// Get - will fetch a tenant by its id.
func (s *Store) Get(ctx context.Context, id string) (tenant.Tenant, error) {
	const q = `
	SELECT *
	FROM tenant
	WHERE id = $1
	`
	var out dbTenant
	if err := s.db.Get(ctx, &out, q, id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return tenant.Tenant{}, core.ErrNotFound
		}
		return tenant.Tenant{}, fmt.Errorf("get: failed to fetch tenant: %w", err)
	}
	return toCoreTenant(out), nil
}

// List - will list all tenants for the provided property id.
func (s *Store) List(ctx context.Context, propertyID string) ([]tenant.Tenant, error) {
	const q = `
	SELECT *
	FROM tenant
	WHERE property_id = $1
	`
	var out []dbTenant
	if err := s.db.List(ctx, &out, q, propertyID); err != nil {
		return nil, fmt.Errorf("list: failed to list tenants: %w", err)
	}
	return toCoreTenants(out), nil
}
