package v1

import (
	"context"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/lenguti/ezuzu/foundation/api"
)

// CreateInvoiceRequest - represents input for creating a new invoice.
type CreateInvoiceRequest struct {
	TenantID string `json:"tenantId"`
	DueDate  string `json:"dueDate"`

	parsedTime time.Time
}

func (cir *CreateInvoiceRequest) validate() *api.ValidationError {
	e := api.NewValidationError()

	if _, err := uuid.Parse(cir.TenantID); err != nil {
		e.Add("tenant id", "is invalid")
	}

	const layout = "2006-01-02"
	tm, err := time.Parse(layout, cir.DueDate)
	if err != nil {
		e.Add("due date", "invalid formet [yyyy-mm-dd]")
	}
	cir.parsedTime = tm

	return e
}

// CreateInvoiceResponse - represents a client create invoice response.
type CreateInvoiceResponse struct {
	Invoice ClientInvoice `json:"invoice"`
}

// CreateInvoice - invoked by POST /v1/managers/:id/properties/invoices.
func (c *Controller) CreateInvoice(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	c.log.Info().Msg("Creating Invoice.")

	var input CreateInvoiceRequest
	if err := api.Decode(r, &input); err != nil {
		c.log.Err(err).Msg("Unable to decode create invoice request.")
		return api.BadRequestError("Invalid input.", err, nil)
	}

	if validated := input.validate(); !validated.IsClean() {
		c.log.Err(validated).Msg("Validation input failed.")
		return api.BadRequestError("Invalid input.", validated, validated.Details())
	}

	mID, err := uuid.Parse(api.PathParam(r, managerIDPathParam))
	if err != nil {
		c.log.Err(err).Msg("Invalid manager id.")
		return api.BadRequestError("Invalid id.", err, nil)
	}

	pID, err := uuid.Parse(api.PathParam(r, propertyIDPathParam))
	if err != nil {
		c.log.Err(err).Msg("Invalid property id.")
		return api.BadRequestError("Invalid id.", err, nil)
	}

	p, err := c.PC.GetProperty(mID, pID)
	if err != nil {
		c.log.Err(err).Msg("Unable to fetch property.")
		return api.InternalServerError("Error.", err, nil)
	}
	c.log.Info().Interface("property", p).Msg("Successfully fetched property.")

	i, err := c.Invoice.Create(ctx, toCoreNewInvoice(input, p.Rent, mID, pID))
	if err != nil {
		c.log.Err(err).Msg("Unable to create invoice.")
		return api.InternalServerError("Error.", err, nil)
	}

	c.log.Info().Msg("Successfully created Invoice.")
	return api.Respond(w, http.StatusCreated, CreateInvoiceResponse{Invoice: toClientInvoice(i)})
}
