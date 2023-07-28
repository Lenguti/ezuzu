package v1

import (
	"github.com/google/uuid"
	"github.com/lenguti/ezuzu/business/core/payment"
)

// ClientPayment - represents a client payment entity.
type ClientPayment struct {
	ID        string  `json:"id"`
	InvoiceID string  `json:"invoiceId"`
	Amount    float64 `json:"amount"`
	CreatedAt int64   `json:"createdAt"`
	UpdatedAt int64   `json:"updatedAt"`
}

func toCoreNewPayment(input CreatePaymentRequest, tID, iID uuid.UUID) payment.NewPayment {
	newPayment := payment.NewPayment{
		TenantID:  tID,
		InvoiceID: iID,
		Amount:    input.Amount,
	}
	return newPayment
}

func toClientPayment(input payment.Payment) ClientPayment {
	return ClientPayment{
		ID:        input.ID.String(),
		InvoiceID: input.InvoiceID.String(),
		Amount:    input.Amount,
		CreatedAt: input.CreatedAt.Unix(),
		UpdatedAt: input.CreatedAt.Unix(),
	}
}
