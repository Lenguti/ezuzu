package v1tests

import (
	"context"

	"github.com/lenguti/ezuzu/business/core/tenant"
)

type mockTenantStore struct {
	tenant.Storer

	createFunc func() error
	updateFunc func() error
	getFunc    func() (tenant.Tenant, error)
}

func (mts *mockTenantStore) Create(ctx context.Context, t tenant.Tenant) error {
	return mts.createFunc()
}

func (mts *mockTenantStore) Update(ctx context.Context, t tenant.Tenant) error {
	return mts.updateFunc()
}

func (mts *mockTenantStore) Get(ctx context.Context, id string) (tenant.Tenant, error) {
	return mts.getFunc()
}
