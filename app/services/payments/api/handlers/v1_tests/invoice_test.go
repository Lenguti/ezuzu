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
	v1 "github.com/lenguti/ezuzu/app/services/payments/api/handlers/v1"
	pv1 "github.com/lenguti/ezuzu/app/services/property/api/handlers/v1"
	"github.com/lenguti/ezuzu/business/core/invoice"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateInvoice(t *testing.T) {
	// Setup.
	ctx := context.Background()
	log := zerolog.New(os.Stdout).With().Timestamp().Logger()
	managerID, propertyID, TenantID := uuid.New(), uuid.New(), uuid.New()
	ctx = httptreemux.AddParamsToContext(ctx, map[string]string{
		"managerId":  managerID.String(),
		"propertyId": propertyID.String(),
	})

	t.Run("success", func(t *testing.T) {
		input := v1.CreateInvoiceRequest{
			TenantID: TenantID.String(),
			DueDate:  "2023-04-15",
		}
		want := v1.ClientInvoice{
			TenantID: TenantID.String(),
			DueDate:  "April 15 2023",
			Amount:   600,
		}
		ctrl := v1.Controller{
			Invoice: invoice.NewCore(&mockInvoiceStore{
				createFunc: func() error {
					return nil
				},
			}, log),
			PC: &mockPropertyClient{
				getPropertyFunc: func() (pv1.ClientProperty, error) {
					return pv1.ClientProperty{
						ID:   propertyID.String(),
						Rent: 600,
					}, nil
				},
			},
		}

		bs, err := json.Marshal(input)
		require.NoError(t, err)

		w := httptest.NewRecorder()
		r, err := http.NewRequestWithContext(ctx, http.MethodPost, "/managers/:managerId/properties/:propertyId/invoices", bytes.NewBuffer(bs))
		require.NoError(t, err)

		// Execute.
		err = ctrl.CreateInvoice(ctx, w, r)

		// Validate.
		require.NoError(t, err)

		var got v1.CreateInvoiceResponse
		err = json.Unmarshal(w.Body.Bytes(), &got)
		require.NoError(t, err)

		assert.Equal(t, want.Amount, got.Invoice.Amount)
		assert.Equal(t, want.DueDate, got.Invoice.DueDate)
		assert.Equal(t, want.TenantID, got.Invoice.TenantID)
	})
}
