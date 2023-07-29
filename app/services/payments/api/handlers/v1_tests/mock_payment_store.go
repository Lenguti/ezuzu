package v1tests

import (
	"context"

	"github.com/lenguti/ezuzu/business/core/payment"
)

type mockPaymentStore struct {
	payment.Storer

	createFunc func() error
	listFunc   func() ([]payment.Payment, error)
}

func (mps *mockPaymentStore) Create(ctx context.Context, p payment.Payment) error {
	return mps.createFunc()
}

func (mps *mockPaymentStore) ListByInvoice(ctx context.Context, invoiceID string) ([]payment.Payment, error) {
	return mps.listFunc()
}
