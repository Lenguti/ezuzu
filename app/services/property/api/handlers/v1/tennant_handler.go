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
	"github.com/lenguti/ezuzu/business/core/tenant"
	"github.com/lenguti/ezuzu/foundation/api"
)

// CreateTenantRequest - represents input for creating a new tenant.
type CreateTenantRequest struct {
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Type        string `json:"type"`
	DateOfBirth string `json:"dateOfBirth"`
	SSN         int    `json:"ssn"`
}

func (ctr *CreateTenantRequest) validate() *api.ValidationError {
	e := api.NewValidationError()

	if ctr.FirstName == "" {
		e.Add("first_name", "is required")
	}

	if ctr.LastName == "" {
		e.Add("last_name", "is required")
	}

	if err := tenant.ParseType(ctr.Type); err != nil {
		e.Add("type", fmt.Sprintf("invalid tenant type, [%s, %s]", tenant.TypePrimary, tenant.TypeSecondary))
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

// CreateTenantResponse - represents a client create tenant response.
type CreateTenantResponse struct {
	Tenant ClientTenant `json:"tenant"`
}

// CreateTenant - invoked by POST /v1/managers/:id/properties/:id/tenants.
func (c *Controller) CreateTenant(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	c.log.Info().Msg("Creating Tenant.")

	var input CreateTenantRequest
	if err := api.Decode(r, &input); err != nil {
		c.log.Err(err).Msg("Unable to decode create tenant request.")
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

	t, err := c.Tenant.Create(ctx, toCoreNewTenant(input), pID)
	if err != nil {
		c.log.Err(err).Msg("Unable to create tenant.")
		return api.InternalServerError("Error.", err, nil)
	}

	c.log.Info().Msg("Successfully created Tenant.")
	return api.Respond(w, http.StatusCreated, CreateTenantResponse{Tenant: toClientTenant(t)})
}

// UpdateTenantRequest - represents input for updating an existing tenant.
type UpdateTenantRequest struct {
	NewPropertyID *string `json:"newPropertyId"`
	Type          *string `json:"type"`
}

func (utr *UpdateTenantRequest) validate() *api.ValidationError {
	e := api.NewValidationError()
	if utr.Type != nil {
		if err := tenant.ParseType(*utr.Type); err != nil {
			e.Add("type", "invalid tenant type")
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

// UpdateTenantResponse - represents a client update tenant response.
type UpdateTenantResponse struct {
	Tenant ClientTenant `json:"tenant"`
}

// UpdateTenant - invoked by PATCH /v1/managers/:id/properties/:id/tenants/:id.
func (c *Controller) UpdateTenant(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	c.log.Info().Msg("Updating Tenant.")

	var input UpdateTenantRequest
	if err := api.Decode(r, &input); err != nil {
		c.log.Err(err).Msg("Unable to decode update tenant request.")
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

	tID, err := uuid.Parse(api.PathParam(r, tenantIDPathParam))
	if err != nil {
		c.log.Err(err).Msg("Invalid tenant id.")
		return api.BadRequestError("Invalid id.", err, nil)
	}

	t, err := c.Tenant.Update(ctx, tID, toCoreUpdateTenant(input))
	if err != nil {
		c.log.Err(err).Msg("Unable to update tenant.")
		if errors.Is(err, core.ErrNotFound) {
			return api.NotFoundError("Item not found.", err, nil)
		}
		return api.InternalServerError("Error.", err, nil)
	}

	c.log.Info().Msg("Successfully updated Property.")
	return api.Respond(w, http.StatusOK, UpdateTenantResponse{Tenant: toClientTenant(t)})
}

// GetTenantResponse - represents a client get tenant response.
type GetTenantResponse struct {
	Tenant ClientTenant `json:"tenant"`
}

// GetTenant - invoked by GET /v1/managers/:id/properties/:id/tenants/:id.
func (c *Controller) GetTenant(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	c.log.Info().Msg("Fetching Tenant.")

	if _, err := uuid.Parse(api.PathParam(r, managerIDPathParam)); err != nil {
		c.log.Err(err).Msg("Invalid manager id.")
		return api.BadRequestError("Invalid id.", err, nil)
	}

	if _, err := uuid.Parse(api.PathParam(r, propertyIDPathParam)); err != nil {
		c.log.Err(err).Msg("Invalid property id.")
		return api.BadRequestError("Invalid id.", err, nil)
	}

	tID, err := uuid.Parse(api.PathParam(r, tenantIDPathParam))
	if err != nil {
		c.log.Err(err).Msg("Invalid tenant id.")
		return api.BadRequestError("Invalid id.", err, nil)
	}

	t, err := c.Tenant.Get(ctx, tID)
	if err != nil {
		c.log.Err(err).Msg("Unable to get tenant.")
		if errors.Is(err, core.ErrNotFound) {
			return api.NotFoundError("Item not found.", err, nil)
		}
		return api.InternalServerError("Error.", err, nil)
	}

	c.log.Info().Msg("Successfully Fetched Tenant.")
	return api.Respond(w, http.StatusOK, GetTenantResponse{Tenant: toClientTenant(t)})
}

// ListTenantsResponse - represents a client list tenant response.
type ListTenantsResponse struct {
	Tenants []ClientTenant `json:"tenants"`
}

// ListTenants - invoked by GET /v1/managers/:id/properties/:id/tenants.
func (c *Controller) ListTenants(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	c.log.Info().Msg("Listing Tenants.")

	if _, err := uuid.Parse(api.PathParam(r, managerIDPathParam)); err != nil {
		c.log.Err(err).Msg("Invalid manager id.")
		return api.BadRequestError("Invalid id.", err, nil)
	}

	pID, err := uuid.Parse(api.PathParam(r, propertyIDPathParam))
	if err != nil {
		c.log.Err(err).Msg("Invalid property id.")
		return api.BadRequestError("Invalid id.", err, nil)
	}

	tts, err := c.Tenant.List(ctx, pID)
	if err != nil {
		c.log.Err(err).Msg("Unable to list tenants.")
		return api.InternalServerError("Error.", err, nil)
	}

	c.log.Info().Msg("Successfully Listed Tenants.")
	return api.Respond(w, http.StatusOK, ListTenantsResponse{Tenants: toClientTenants(tts)})
}
