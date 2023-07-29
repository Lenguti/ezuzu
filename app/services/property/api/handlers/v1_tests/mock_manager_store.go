package v1tests

import (
	"context"

	"github.com/lenguti/ezuzu/business/core/manager"
)

type mockManagerStore struct {
	manager.Storer

	createFunc func() error
}

func (mms *mockManagerStore) Create(ctx context.Context, m manager.Manager) error {
	return mms.createFunc()
}
