package v1

import (
	"context"
	"net/http"

	"github.com/lenguti/ezuzu/foundation/api"
)

// CreateManagerRequest - represents input for creating a new property manager.
type CreateManagerRequest struct {
	Entity string `json:"entity"`
}

func (cmr *CreateManagerRequest) validate() *api.ValidationError {
	e := api.NewValidationError()
	if cmr.Entity == "" {
		e.Add("entity", "is required")
	}

	return e
}

// CreateManagerResponse - represents a client create property manager response.
type CreateManagerResponse struct {
	Manager ClientManager `json:"manager"`
}

// CreateManager - invoked by POST /v1/managers.
func (c *Controller) CreateManager(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	c.log.Info().Msg("Creating Manager.")

	var input CreateManagerRequest
	if err := api.Decode(r, &input); err != nil {
		c.log.Err(err).Msg("Unable to decode create manager request.")
		return api.BadRequestError("Invalid input.", err, nil)
	}

	if validated := input.validate(); !validated.IsClean() {
		c.log.Err(validated).Msg("Validation input failed.")
		return api.BadRequestError("Invalid input.", validated, validated.Details())
	}

	m, err := c.Manager.Create(ctx, toCoreNewManager(input))
	if err != nil {
		c.log.Err(err).Msg("Unable to create manager.")
		return api.InternalServerError("Error.", err, nil)
	}

	c.log.Info().Msg("Successfully created Manager.")
	return api.Respond(w, http.StatusCreated, CreateManagerResponse{Manager: toClientManager(m)})
}
