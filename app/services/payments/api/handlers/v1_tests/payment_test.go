package v1tests

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/dimfeld/httptreemux"
	"github.com/google/uuid"
	v1 "github.com/lenguti/ezuzu/app/services/payments/api/handlers/v1"
	"github.com/lenguti/ezuzu/business/core/invoice"
	"github.com/lenguti/ezuzu/business/core/payment"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreatePayment(t *testing.T) {
	ctx := context.Background()
	log := zerolog.New(os.Stdout).With().Timestamp().Logger()
	tenantID, invoiceID := uuid.New(), uuid.New()
	ctx = httptreemux.AddParamsToContext(ctx, map[string]string{
		"tenantId":  tenantID.String(),
		"invoiceId": invoiceID.String(),
	})

	t.Run("success", func(t *testing.T) {
		// Setup.
		input := v1.CreatePaymentRequest{
			Amount: 500,
		}
		want := v1.ClientPayment{
			InvoiceID: invoiceID.String(),
			Amount:    500,
		}
		ctrl := v1.Controller{
			Payment: payment.NewCore(&mockPaymentStore{
				listFunc: func() ([]payment.Payment, error) {
					return nil, nil
				},
				createFunc: func() error {
					return nil
				},
			},
				invoice.NewCore(&mockInvoiceStore{
					getFunc: func() (invoice.Invoice, error) {
						return invoice.Invoice{
							Amount:  500,
							DueDate: time.Now().Add(24 * time.Hour),
						}, nil
					},
				}, log), log),
		}

		bs, err := json.Marshal(input)
		require.NoError(t, err)

		w := httptest.NewRecorder()
		r, err := http.NewRequestWithContext(ctx, http.MethodPost, "/tenants/:tenantId/invoices/:invoiceId/payments", bytes.NewBuffer(bs))
		require.NoError(t, err)

		// Execute.
		err = ctrl.CreatePayment(ctx, w, r)

		// Validate.
		require.NoError(t, err)

		var got v1.CreatePaymentResponse
		err = json.Unmarshal(w.Body.Bytes(), &got)
		require.NoError(t, err)

		assert.Equal(t, want.InvoiceID, got.Payment.InvoiceID)
		assert.Equal(t, want.Amount, got.Payment.Amount)
	})
}
