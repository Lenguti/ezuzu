package v1tests

import (
	"context"

	"github.com/lenguti/ezuzu/business/core/property"
)

type mockPropertyStore struct {
	property.Storer

	createFunc func() error
	getFunc    func() (property.Property, error)
	updateFunc func() error
}

func (mps *mockPropertyStore) Create(ctx context.Context, p property.Property) error {
	return mps.createFunc()
}

func (mps *mockPropertyStore) Get(ctx context.Context, id string) (property.Property, error) {
	return mps.getFunc()
}

func (mps *mockPropertyStore) Update(ctx context.Context, p property.Property) error {
	return mps.updateFunc()
}
