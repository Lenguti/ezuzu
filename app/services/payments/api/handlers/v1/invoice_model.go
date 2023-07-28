package v1

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/lenguti/ezuzu/business/core/invoice"
)

// ClientInvoice - represents a client invoice entity.
type ClientInvoice struct {
	ID        string  `json:"id"`
	TenantID  string  `json:"tenantId"`
	Amount    float64 `json:"amount"`
	DueDate   string  `json:"dueDate"`
	CreatedAt int64   `json:"createdAt"`
	UpdatedAt int64   `json:"updatedAt"`
}

func toCoreNewInvoice(input CreateInvoiceRequest, amount float64, mID, pID uuid.UUID) invoice.NewInvoice {
	newInvoice := invoice.NewInvoice{
		ManagerID:  mID,
		PropertyID: pID,
		TenantID:   uuid.MustParse(input.TenantID),
		Amount:     amount,
		DueDate:    input.parsedTime,
	}
	return newInvoice
}

func toClientInvoice(input invoice.Invoice) ClientInvoice {
	return ClientInvoice{
		ID:        input.ID.String(),
		TenantID:  input.TenantID.String(),
		Amount:    input.Amount,
		DueDate:   fmt.Sprintf("%v %v %v", input.DueDate.Month(), input.DueDate.Day(), input.DueDate.Year()),
		CreatedAt: input.CreatedAt.Unix(),
		UpdatedAt: input.UpdatedAt.Unix(),
	}
}
