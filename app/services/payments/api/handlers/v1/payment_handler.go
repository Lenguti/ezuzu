package v1

import (
	"context"
	"errors"
	"net/http"

	"github.com/google/uuid"
	"github.com/lenguti/ezuzu/business/core"
	"github.com/lenguti/ezuzu/foundation/api"
)

// CreatePaymentRequest - represents input for creating a new payment.
type CreatePaymentRequest struct {
	Amount float64 `json:"amount"`
}

func (cpr *CreatePaymentRequest) validate() *api.ValidationError {
	e := api.NewValidationError()

	if cpr.Amount <= 0 {
		e.Add("amount", "invalid amount, must pay more than 0")
	}

	return e
}

// CreatePaymentResponse - represents a client create payment response.
type CreatePaymentResponse struct {
	Payment ClientPayment `json:"payment"`
}

// CreatePayment - invoked by POST /v1/tenants/:id/invoices/:id/payments.
func (c *Controller) CreatePayment(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	c.log.Info().Msg("Creating Payment.")

	var input CreatePaymentRequest
	if err := api.Decode(r, &input); err != nil {
		c.log.Err(err).Msg("Unable to decode create payment request.")
		return api.BadRequestError("Invalid input.", err, nil)
	}

	if validated := input.validate(); !validated.IsClean() {
		c.log.Err(validated).Msg("Validation input failed.")
		return api.BadRequestError("Invalid input.", validated, validated.Details())
	}

	tID, err := uuid.Parse(api.PathParam(r, tenantIDPathParam))
	if err != nil {
		c.log.Err(err).Msg("Invalid tenant id.")
		return api.BadRequestError("Invalid id.", err, nil)
	}

	iID, err := uuid.Parse(api.PathParam(r, invoiceIDPathParam))
	if err != nil {
		c.log.Err(err).Msg("Invalid invoice id.")
		return api.BadRequestError("Invalid id.", err, nil)
	}

	p, err := c.Payment.Create(ctx, toCoreNewPayment(input, tID, iID))
	if err != nil {
		c.log.Err(err).Msg("Unable to create payment.")
		if errors.Is(err, core.ErrPastDuePayment) || errors.Is(err, core.ErrPaymentConflict) {
			return api.BadRequestError("Invalid payment request.", err, nil)
		}
		return api.InternalServerError("Error.", err, nil)
	}

	c.log.Info().Msg("Successfully created Payment.")
	return api.Respond(w, http.StatusCreated, CreatePaymentResponse{Payment: toClientPayment(p)})
}
