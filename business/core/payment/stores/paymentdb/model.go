package paymentdb

import (
	"time"

	"github.com/google/uuid"
	"github.com/lenguti/ezuzu/business/core/payment"
)

type dbPayment struct {
	ID        string  `db:"id"`
	TenantID  string  `db:"tenant_id"`
	InvoiceID string  `db:"invoice_id"`
	Amount    float64 `db:"amount"`
	CreatedAt int64   `db:"created_at"`
	UpdatedAt int64   `db:"updated_at"`
}

func toDBPayment(p payment.Payment) dbPayment {
	return dbPayment{
		ID:        p.ID.String(),
		TenantID:  p.TenantID.String(),
		InvoiceID: p.InvoiceID.String(),
		Amount:    p.Amount,
		CreatedAt: p.CreatedAt.Unix(),
		UpdatedAt: p.UpdatedAt.Unix(),
	}
}

func toCorePayments(dbps []dbPayment) []payment.Payment {
	ps := make([]payment.Payment, 0, len(dbps))
	for _, dbp := range dbps {
		ps = append(ps, toCorePayment(dbp))
	}
	return ps
}

func toCorePayment(dbp dbPayment) payment.Payment {
	return payment.Payment{
		ID:        uuid.MustParse(dbp.ID),
		TenantID:  uuid.MustParse(dbp.TenantID),
		InvoiceID: uuid.MustParse(dbp.InvoiceID),
		Amount:    dbp.Amount,
		CreatedAt: time.Unix(dbp.CreatedAt, 0),
		UpdatedAt: time.Unix(dbp.UpdatedAt, 0),
	}
}
