package v1tests

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/dimfeld/httptreemux"
	"github.com/google/uuid"
	v1 "github.com/lenguti/ezuzu/app/services/property/api/handlers/v1"
	"github.com/lenguti/ezuzu/business/core/property"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateProperty(t *testing.T) {
	// Setup.
	ctx := context.Background()
	log := zerolog.New(os.Stdout).With().Timestamp().Logger()
	managerID := uuid.New()
	ctx = httptreemux.AddParamsToContext(ctx, map[string]string{
		"managerId": managerID.String(),
	})

	t.Run("success", func(t *testing.T) {
		unit := 10
		input := v1.CreatePropertyRequest{
			Street:     "210 NW 15 Ave",
			City:       "Plantation",
			State:      "FL",
			PostalCode: "33324",
			Name:       "Cool Biz Apartments",
			Rent:       1250,
			Type:       "APARTMENT",
			UnitNumber: &unit,
		}
		want := v1.ClientProperty{
			Address:    "210 NW 15 Ave, Plantation, FL 33324",
			Name:       "Cool Biz Apartments",
			Rent:       1250,
			Type:       "APARTMENT",
			UnitNumber: &unit,
		}
		ctrl := v1.Controller{
			Property: property.NewCore(&mockPropertyStore{
				createFunc: func() error {
					return nil
				},
			}, log),
		}

		bs, err := json.Marshal(input)
		require.NoError(t, err)

		w := httptest.NewRecorder()
		r, err := http.NewRequestWithContext(ctx, http.MethodPost, "/managers/:managerId/properties", bytes.NewBuffer(bs))
		require.NoError(t, err)

		// Execute.
		err = ctrl.CreateProperty(ctx, w, r)

		// Validate.
		require.NoError(t, err)

		var got v1.CreatePropertyResponse
		err = json.Unmarshal(w.Body.Bytes(), &got)
		require.NoError(t, err)

		assert.Equal(t, want.Address, got.Property.Address)
		assert.Equal(t, want.Name, got.Property.Name)
		assert.Equal(t, want.Rent, got.Property.Rent)
		assert.Equal(t, want.Type, got.Property.Type)
		assert.Equal(t, want.UnitNumber, got.Property.UnitNumber)
	})
}

func TestUpdateProperty(t *testing.T) {
	// Setup.
	ctx := context.Background()
	log := zerolog.New(os.Stdout).With().Timestamp().Logger()
	managerID, propertyID := uuid.New(), uuid.New()
	ctx = httptreemux.AddParamsToContext(ctx, map[string]string{
		"managerId":  managerID.String(),
		"propertyId": propertyID.String(),
	})

	t.Run("success", func(t *testing.T) {
		newName := "New Cool Biz Apartments"
		newRent := float64(1000)
		input := v1.UpdatePropertyRequest{
			Name: &newName,
			Rent: &newRent,
		}
		want := v1.ClientProperty{
			Address: "210 NW 15 Ave, Plantation, FL 33324",
			Name:    "New Cool Biz Apartments",
			Rent:    1000,
			Type:    "APARTMENT",
		}
		ctrl := v1.Controller{
			Property: property.NewCore(&mockPropertyStore{
				getFunc: func() (property.Property, error) {
					return property.Property{
						Address: "210 NW 15 Ave, Plantation, FL 33324",
						Type:    "APARTMENT",
					}, nil
				},
				updateFunc: func() error {
					return nil
				},
			}, log),
		}

		bs, err := json.Marshal(input)
		require.NoError(t, err)

		w := httptest.NewRecorder()
		r, err := http.NewRequestWithContext(ctx, http.MethodPatch, "/managers/:managerId/properties/:propertyId", bytes.NewBuffer(bs))
		require.NoError(t, err)

		// Execute.
		err = ctrl.UpdateProperty(ctx, w, r)

		// Validate.
		require.NoError(t, err)

		var got v1.UpdatePropertyResponse
		err = json.Unmarshal(w.Body.Bytes(), &got)
		require.NoError(t, err)

		assert.Equal(t, want.Address, got.Property.Address)
		assert.Equal(t, want.Name, got.Property.Name)
		assert.Equal(t, want.Rent, got.Property.Rent)
		assert.Equal(t, want.Type, got.Property.Type)
	})
}
