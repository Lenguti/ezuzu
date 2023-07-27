package tennantdb

import (
	"time"

	"github.com/google/uuid"
	"github.com/lenguti/ezuzu/business/core/tennant"
)

type dbTennant struct {
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

func toDBTennant(t tennant.Tennant) dbTennant {
	return dbTennant{
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

func toCoreTennants(dbtts []dbTennant) []tennant.Tennant {
	tts := make([]tennant.Tennant, 0, len(dbtts))
	for _, t := range dbtts {
		tts = append(tts, toCoreTennant(t))
	}
	return tts
}

func toCoreTennant(dbt dbTennant) tennant.Tennant {
	return tennant.Tennant{
		ID:          uuid.MustParse(dbt.ID),
		PropertyID:  uuid.MustParse(dbt.PropertyID),
		FirstName:   dbt.FirstName,
		LastName:    dbt.LastName,
		DateOfBirth: dbt.DateOfBirth,
		SSN:         dbt.SSN,
		Type:        tennant.Type(dbt.Type),
		CreatedAt:   time.Unix(dbt.CreatedAt, 0),
		UpdatedAt:   time.Unix(dbt.UpdatedAt, 0),
	}
}
