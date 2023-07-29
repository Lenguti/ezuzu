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
	"github.com/lenguti/ezuzu/business/core/tenant"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateTenant(t *testing.T) {
	// Setup.
	ctx := context.Background()
	log := zerolog.New(os.Stdout).With().Timestamp().Logger()
	managerID, propertyID := uuid.New(), uuid.New()
	ctx = httptreemux.AddParamsToContext(ctx, map[string]string{
		"managerId":  managerID.String(),
		"propertyId": propertyID.String(),
	})

	t.Run("success", func(t *testing.T) {
		input := v1.CreateTenantRequest{
			FirstName:   "Johnnie",
			LastName:    "Boi",
			Type:        "PRIMARY",
			DateOfBirth: "1987-04-20",
			SSN:         986543862,
		}
		want := v1.ClientTenant{
			FirstName:   "Johnnie",
			LastName:    "Boi",
			Type:        "PRIMARY",
			DateOfBirth: "1987-04-20",
			SSN:         "#########",
		}
		ctrl := v1.Controller{
			Tenant: tenant.NewCore(&mockTenantStore{
				createFunc: func() error {
					return nil
				},
			},
				property.NewCore(&mockPropertyStore{}, log),
				log),
		}

		bs, err := json.Marshal(input)
		require.NoError(t, err)

		w := httptest.NewRecorder()
		r, err := http.NewRequestWithContext(ctx, http.MethodPost, "/managers/:managerId/properties/:propertyId/tenants", bytes.NewBuffer(bs))
		require.NoError(t, err)

		// Execute.
		err = ctrl.CreateTenant(ctx, w, r)

		// Validate.
		require.NoError(t, err)

		var got v1.CreateTenantResponse
		err = json.Unmarshal(w.Body.Bytes(), &got)
		require.NoError(t, err)

		assert.Equal(t, want.FirstName, got.Tenant.FirstName)
		assert.Equal(t, want.LastName, got.Tenant.LastName)
		assert.Equal(t, want.Type, got.Tenant.Type)
		assert.Equal(t, want.DateOfBirth, got.Tenant.DateOfBirth)
		assert.Equal(t, want.SSN, got.Tenant.SSN)
	})
}

func TestUpdateTenant(t *testing.T) {
	// Setup.
	ctx := context.Background()
	log := zerolog.New(os.Stdout).With().Timestamp().Logger()
	managerID, propertyID, tennantID := uuid.New(), uuid.New(), uuid.New()
	ctx = httptreemux.AddParamsToContext(ctx, map[string]string{
		"managerId":  managerID.String(),
		"propertyId": propertyID.String(),
		"tenantId":   tennantID.String(),
	})

	t.Run("success", func(t *testing.T) {
		newPropertyID := uuid.New().String()
		newType := "SECONDARY"
		input := v1.UpdateTenantRequest{
			NewPropertyID: &newPropertyID,
			Type:          &newType,
		}
		want := v1.ClientTenant{
			Type: newType,
		}
		ctrl := v1.Controller{
			Tenant: tenant.NewCore(&mockTenantStore{
				getFunc: func() (tenant.Tenant, error) {
					return tenant.Tenant{
						Type: "PRIMARY",
					}, nil
				},
				updateFunc: func() error {
					return nil
				}},
				property.NewCore(&mockPropertyStore{
					getFunc: func() (property.Property, error) {
						return property.Property{
							ID: propertyID,
						}, nil
					},
				}, log),
				log),
		}

		bs, err := json.Marshal(input)
		require.NoError(t, err)

		w := httptest.NewRecorder()
		r, err := http.NewRequestWithContext(ctx, http.MethodPatch, "/managers/:managerId/properties/:propertyId/tenants/:tenantId", bytes.NewBuffer(bs))
		require.NoError(t, err)

		// Execute.
		err = ctrl.UpdateTenant(ctx, w, r)

		// Validate.
		require.NoError(t, err)

		var got v1.UpdateTenantResponse
		err = json.Unmarshal(w.Body.Bytes(), &got)
		require.NoError(t, err)

		assert.Equal(t, want.Type, got.Tenant.Type)
	})
}
