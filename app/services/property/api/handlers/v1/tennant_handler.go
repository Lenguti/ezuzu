package v1

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"strconv"

	"github.com/google/uuid"
	"github.com/lenguti/ezuzu/business/core"
	"github.com/lenguti/ezuzu/business/core/tennant"
	"github.com/lenguti/ezuzu/foundation/api"
)

// CreateTennantRequest - represents input for creating a new tennant.
type CreateTennantRequest struct {
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Type        string `json:"type"`
	DateOfBirth string `json:"dateOfBirth"`
	SSN         int    `json:"ssn"`
}

func (ctr *CreateTennantRequest) validate() *api.ValidationError {
	e := api.NewValidationError()

	if ctr.FirstName == "" {
		e.Add("first_name", "is required")
	}

	if ctr.LastName == "" {
		e.Add("last_name", "is required")
	}

	if err := tennant.ParseType(ctr.Type); err != nil {
		e.Add("type", fmt.Sprintf("invalid tennant type, [%s, %s]", tennant.TypePrimary, tennant.TypeSecondary))
	}

	re := regexp.MustCompile(`\d{4}-\d{2}-\d{2}`)
	if !re.MatchString(ctr.DateOfBirth) {
		e.Add("type", "invalid date of birth format, [yyyy-mm-dd]")
	}

	if len(strconv.Itoa(ctr.SSN)) != 9 {
		e.Add("ssn", "invalid ssn length")
	}
	return e
}

// CreateTennantResponse - represents a client create tennant response.
type CreateTennantResponse struct {
	Tennant ClientTennant `json:"tennant"`
}

// CreateTennant - invoked by POST /v1/managers/:id/properties/:id/tennants.
func (c *Controller) CreateTennant(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	c.log.Info().Msg("Creating Tennant.")

	var input CreateTennantRequest
	if err := api.Decode(r, &input); err != nil {
		c.log.Err(err).Msg("Unable to decode create tennant request.")
		return api.BadRequestError("Invalid input.", err, nil)
	}

	if validated := input.validate(); !validated.IsClean() {
		c.log.Err(validated).Msg("Validation input failed.")
		return api.BadRequestError("Invalid input.", validated, validated.Details())
	}

	if _, err := uuid.Parse(api.PathParam(r, managerIDPathParam)); err != nil {
		c.log.Err(err).Msg("Invalid manager id.")
		return api.BadRequestError("Invalid id.", err, nil)
	}

	pID, err := uuid.Parse(api.PathParam(r, propertyIDPathParam))
	if err != nil {
		c.log.Err(err).Msg("Invalid property id.")
		return api.BadRequestError("Invalid id.", err, nil)
	}

	t, err := c.Tennant.Create(ctx, toCoreNewTennant(input), pID)
	if err != nil {
		c.log.Err(err).Msg("Unable to create tennant.")
		return api.InternalServerError("Error.", err, nil)
	}

	c.log.Info().Msg("Successfully created Tennant.")
	return api.Respond(w, http.StatusCreated, CreateTennantResponse{Tennant: toClientTennant(t)})
}

// UpdateTennantRequest - represents input for updating an existing tennant.
type UpdateTennantRequest struct {
	NewPropertyID *string `json:"newPropertyId"`
	Type          *string `json:"type"`
}

func (utr *UpdateTennantRequest) validate() *api.ValidationError {
	e := api.NewValidationError()
	if utr.Type != nil {
		if err := tennant.ParseType(*utr.Type); err != nil {
			e.Add("type", "invalid tennant type")
		}
	}

	if utr.NewPropertyID != nil {
		if _, err := uuid.Parse(*utr.NewPropertyID); err != nil {
			e.Add("property id", "is invalid")
		}
	}

	if utr.Type == nil && utr.NewPropertyID == nil {
		e.Add("input", "input must be provided")
	}
	return e
}

// UpdateTennantResponse - represents a client update tennant response.
type UpdateTennantResponse struct {
	Tennant ClientTennant `json:"tennant"`
}

// UpdateTennant - invoked by PATCH /v1/managers/:id/properties/:id/tennants/:id.
func (c *Controller) UpdateTennant(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	c.log.Info().Msg("Updating Tennant.")

	var input UpdateTennantRequest
	if err := api.Decode(r, &input); err != nil {
		c.log.Err(err).Msg("Unable to decode update tennant request.")
		return api.BadRequestError("Invalid input.", err, nil)
	}

	if validated := input.validate(); !validated.IsClean() {
		c.log.Err(validated).Msg("Validation input failed.")
		return api.BadRequestError("Invalid input.", validated, validated.Details())
	}

	if _, err := uuid.Parse(api.PathParam(r, managerIDPathParam)); err != nil {
		c.log.Err(err).Msg("Invalid manager id.")
		return api.BadRequestError("Invalid id.", err, nil)
	}

	if _, err := uuid.Parse(api.PathParam(r, propertyIDPathParam)); err != nil {
		c.log.Err(err).Msg("Invalid property id.")
		return api.BadRequestError("Invalid id.", err, nil)
	}

	tID, err := uuid.Parse(api.PathParam(r, tennantIDPathParam))
	if err != nil {
		c.log.Err(err).Msg("Invalid tennant id.")
		return api.BadRequestError("Invalid id.", err, nil)
	}

	t, err := c.Tennant.Update(ctx, tID, toCoreUpdateTennant(input))
	if err != nil {
		c.log.Err(err).Msg("Unable to update tennant.")
		if errors.Is(err, core.ErrNotFound) {
			return api.NotFoundError("Item not found.", err, nil)
		}
		return api.InternalServerError("Error.", err, nil)
	}

	c.log.Info().Msg("Successfully updated Property.")
	return api.Respond(w, http.StatusOK, UpdateTennantResponse{Tennant: toClientTennant(t)})
}

// GetTennantResponse - represents a client get tennant response.
type GetTennantResponse struct {
	Tennant ClientTennant `json:"tennant"`
}

// GetTennant - invoked by GET /v1/managers/:id/properties/:id/tennants/:id.
func (c *Controller) GetTennant(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	c.log.Info().Msg("Fetching Tennant.")

	if _, err := uuid.Parse(api.PathParam(r, managerIDPathParam)); err != nil {
		c.log.Err(err).Msg("Invalid manager id.")
		return api.BadRequestError("Invalid id.", err, nil)
	}

	if _, err := uuid.Parse(api.PathParam(r, propertyIDPathParam)); err != nil {
		c.log.Err(err).Msg("Invalid property id.")
		return api.BadRequestError("Invalid id.", err, nil)
	}

	tID, err := uuid.Parse(api.PathParam(r, tennantIDPathParam))
	if err != nil {
		c.log.Err(err).Msg("Invalid tennant id.")
		return api.BadRequestError("Invalid id.", err, nil)
	}

	t, err := c.Tennant.Get(ctx, tID)
	if err != nil {
		c.log.Err(err).Msg("Unable to get tennant.")
		if errors.Is(err, core.ErrNotFound) {
			return api.NotFoundError("Item not found.", err, nil)
		}
		return api.InternalServerError("Error.", err, nil)
	}

	c.log.Info().Msg("Successfully Fetched Tennant.")
	return api.Respond(w, http.StatusOK, GetTennantResponse{Tennant: toClientTennant(t)})
}

// ListTennantsResponse - represents a client list tennant response.
type ListTennantsResponse struct {
	Tennants []ClientTennant `json:"tennants"`
}

// ListTennants - invoked by GET /v1/managers/:id/properties/:id/tennants.
func (c *Controller) ListTennants(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	c.log.Info().Msg("Listing Tennants.")

	if _, err := uuid.Parse(api.PathParam(r, managerIDPathParam)); err != nil {
		c.log.Err(err).Msg("Invalid manager id.")
		return api.BadRequestError("Invalid id.", err, nil)
	}

	pID, err := uuid.Parse(api.PathParam(r, propertyIDPathParam))
	if err != nil {
		c.log.Err(err).Msg("Invalid property id.")
		return api.BadRequestError("Invalid id.", err, nil)
	}

	tts, err := c.Tennant.List(ctx, pID)
	if err != nil {
		c.log.Err(err).Msg("Unable to list tennants.")
		return api.InternalServerError("Error.", err, nil)
	}

	c.log.Info().Msg("Successfully Listed Tennants.")
	return api.Respond(w, http.StatusOK, ListTennantsResponse{Tennants: toClientTennants(tts)})
}
