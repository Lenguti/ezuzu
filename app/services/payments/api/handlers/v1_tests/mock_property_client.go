package v1tests

import (
	"github.com/google/uuid"
	v1 "github.com/lenguti/ezuzu/app/services/property/api/handlers/v1"
	"github.com/lenguti/ezuzu/app/services/property/api/handlers/v1/client"
)

type mockPropertyClient struct {
	client.IProperty

	getPropertyFunc func() (v1.ClientProperty, error)
}

func (mpc *mockPropertyClient) GetProperty(managerID, propertyID uuid.UUID) (v1.ClientProperty, error) {
	return mpc.getPropertyFunc()
}
