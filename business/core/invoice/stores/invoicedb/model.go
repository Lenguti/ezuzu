package invoicedb

import (
	"time"

	"github.com/google/uuid"
	"github.com/lenguti/ezuzu/business/core/invoice"
)

type dbInvoice struct {
	ID         string  `db:"id"`
	ManagerID  string  `db:"manager_id"`
	PropertyID string  `db:"property_id"`
	TenantID   string  `db:"tenant_id"`
	Amount     float64 `db:"amount"`
	DueDate    int64   `db:"due_date"`
	UnitNumber *int    `db:"unit_number"`
	CreatedAt  int64   `db:"created_at"`
	UpdatedAt  int64   `db:"updated_at"`
}

func toDBInvoice(i invoice.Invoice) dbInvoice {
	return dbInvoice{
		ID:         i.ID.String(),
		ManagerID:  i.ManagerID.String(),
		PropertyID: i.PropertyID.String(),
		TenantID:   i.TenantID.String(),
		Amount:     i.Amount,
		DueDate:    i.DueDate.Unix(),
		CreatedAt:  i.CreatedAt.Unix(),
		UpdatedAt:  i.UpdatedAt.Unix(),
	}
}

func toCoreInvoice(dbi dbInvoice) invoice.Invoice {
	return invoice.Invoice{
		ID:         uuid.MustParse(dbi.ID),
		ManagerID:  uuid.MustParse(dbi.ManagerID),
		PropertyID: uuid.MustParse(dbi.PropertyID),
		TenantID:   uuid.MustParse(dbi.TenantID),
		Amount:     dbi.Amount,
		DueDate:    time.Unix(dbi.DueDate, 0),
		CreatedAt:  time.Unix(dbi.CreatedAt, 0),
		UpdatedAt:  time.Unix(dbi.UpdatedAt, 0),
	}
}
