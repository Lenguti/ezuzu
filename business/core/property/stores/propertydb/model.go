package propertydb

import (
	"time"

	"github.com/google/uuid"
	"github.com/lenguti/ezuzu/business/core/property"
)

type dbProperty struct {
	ID         string  `db:"id"`
	ManagerID  string  `db:"manager_id"`
	Address    string  `db:"address"`
	Name       string  `db:"name"`
	Rent       float64 `db:"rent"`
	Type       string  `db:"type"`
	UnitNumber *int    `db:"unit_number"`
	CreatedAt  int64   `db:"created_at"`
	UpdatedAt  int64   `db:"updated_at"`
}

func toDBProperty(p property.Property) dbProperty {
	return dbProperty{
		ID:         p.ID.String(),
		ManagerID:  p.ManagerID.String(),
		Address:    p.Address,
		Name:       p.Name,
		Rent:       p.Rent,
		Type:       p.Type.String(),
		UnitNumber: p.UnitNumber,
		CreatedAt:  p.CreatedAt.Unix(),
		UpdatedAt:  p.UpdatedAt.Unix(),
	}
}

func toCoreProperties(dbpps []dbProperty) []property.Property {
	pps := make([]property.Property, 0, len(dbpps))
	for _, dbp := range dbpps {
		pps = append(pps, toCoreProperty(dbp))
	}
	return pps
}

func toCoreProperty(dbp dbProperty) property.Property {
	return property.Property{
		ID:         uuid.MustParse(dbp.ID),
		ManagerID:  uuid.MustParse(dbp.ManagerID),
		Address:    dbp.Address,
		Name:       dbp.Name,
		Rent:       dbp.Rent,
		Type:       property.Type(dbp.Type),
		UnitNumber: dbp.UnitNumber,
		CreatedAt:  time.Unix(dbp.CreatedAt, 0),
		UpdatedAt:  time.Unix(dbp.UpdatedAt, 0),
	}
}
