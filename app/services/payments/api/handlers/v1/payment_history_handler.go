package v1

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/lenguti/ezuzu/foundation/api"
)

// ListPaymentHistoryResponse - represents a client list payment history response.
type ListPaymentHistoryResponse struct {
	PaymentHistory []ClientPaymentHistory `json:"payment_history"`
}

// ListPaymentHistory - invoked by GET /v1/tenants/:id/payments/history.
func (c *Controller) ListPaymentHistory(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	c.log.Info().Msg("Listing Payment History.")

	tID, err := uuid.Parse(api.PathParam(r, tenantIDPathParam))
	if err != nil {
		c.log.Err(err).Msg("Invalid tenant id.")
		return api.BadRequestError("Invalid id.", err, nil)
	}

	phs, err := c.PaymentHistory.List(ctx, tID)
	if err != nil {
		c.log.Err(err).Msg("Unable to list payment history.")
		return api.InternalServerError("Error.", err, nil)
	}

	c.log.Info().Msg("Successfully Listed Payment History.")
	return api.Respond(w, http.StatusOK, ListPaymentHistoryResponse{PaymentHistory: toClientPaymentHistories(phs)})
}
