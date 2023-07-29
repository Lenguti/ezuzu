package v1tests

import (
	"context"

	"github.com/lenguti/ezuzu/business/core/invoice"
)

type mockInvoiceStore struct {
	invoice.Storer

	createFunc func() error
	getFunc    func() (invoice.Invoice, error)
}

func (mis *mockInvoiceStore) Create(ctx context.Context, i invoice.Invoice) error {
	return mis.createFunc()
}

func (mis *mockInvoiceStore) Get(ctx context.Context, id string) (invoice.Invoice, error) {
	return mis.getFunc()
}
