package tenantdb

import (
	"time"

	"github.com/google/uuid"
	"github.com/lenguti/ezuzu/business/core/tenant"
)

type dbTenant struct {
	ID          string `db:"id"`
	PropertyID  string `db:"property_id"`
	FirstName   string `db:"first_name"`
	LastName    string `db:"last_name"`
	DateOfBirth string `db:"date_of_birth"`
	SSN         int    `db:"ssn"`
	Type        string `db:"type"`
	CreatedAt   int64  `db:"created_at"`
	UpdatedAt   int64  `db:"updated_at"`
}

func toDBTenant(t tenant.Tenant) dbTenant {
	return dbTenant{
		ID:          t.ID.String(),
		PropertyID:  t.PropertyID.String(),
		FirstName:   t.FirstName,
		LastName:    t.LastName,
		DateOfBirth: t.DateOfBirth,
		SSN:         t.SSN,
		Type:        t.Type.String(),
		CreatedAt:   t.CreatedAt.Unix(),
		UpdatedAt:   t.UpdatedAt.Unix(),
	}
}

func toCoreTenants(dbtts []dbTenant) []tenant.Tenant {
	tts := make([]tenant.Tenant, 0, len(dbtts))
	for _, t := range dbtts {
		tts = append(tts, toCoreTenant(t))
	}
	return tts
}

func toCoreTenant(dbt dbTenant) tenant.Tenant {
	return tenant.Tenant{
		ID:          uuid.MustParse(dbt.ID),
		PropertyID:  uuid.MustParse(dbt.PropertyID),
		FirstName:   dbt.FirstName,
		LastName:    dbt.LastName,
		DateOfBirth: dbt.DateOfBirth,
		SSN:         dbt.SSN,
		Type:        tenant.Type(dbt.Type),
		CreatedAt:   time.Unix(dbt.CreatedAt, 0),
		UpdatedAt:   time.Unix(dbt.UpdatedAt, 0),
	}
}
