package v1

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/Boostport/address"
	"github.com/google/uuid"
	"github.com/hashicorp/go-multierror"
	"github.com/lenguti/ezuzu/business/core"
	"github.com/lenguti/ezuzu/business/core/property"
	"github.com/lenguti/ezuzu/foundation/api"
)

// CreatePropertyRequest - represents input for creating a new property.
type CreatePropertyRequest struct {
	Street     string  `json:"street"`
	City       string  `json:"city"`
	State      string  `json:"state"`
	PostalCode string  `json:"postalCode"`
	Name       string  `json:"name"`
	Rent       float64 `json:"rent"`
	Type       string  `json:"type"`
	UnitNumber *int    `json:"unitNumber"`

	formattedAddress string
}

func (cpr *CreatePropertyRequest) validate() *api.ValidationError {
	e := api.NewValidationError()

	addr, err := address.NewValid(
		address.WithStreetAddress([]string{cpr.Street}),
		address.WithLocality(cpr.City),
		address.WithAdministrativeArea(cpr.State),
		address.WithPostCode(cpr.PostalCode),
	)
	if err != nil {
		if merr, ok := errors.Unwrap(err).(*multierror.Error); ok {
			for _, subErr := range merr.Errors {
				switch {
				case errors.Is(subErr, address.ErrInvalidAdministrativeArea):
					e.Add("state", "invalid state")
				case errors.Is(subErr, address.ErrInvalidPostCode):
					e.Add("postal code", "invalid postal code")
				case errors.Is(subErr, address.ErrInvalidLocality):
					e.Add("city", "invalid city")
				}
			}
		}
	}
	cpr.formattedAddress = fmt.Sprintf("%s, %s, %s %s", addr.StreetAddress[0], addr.Locality, addr.AdministrativeArea, addr.PostCode)

	if cpr.Name == "" {
		e.Add("name", "is required")
	}

	if cpr.Rent <= 0 {
		e.Add("rent", "needs to be greater than 0")
	}

	if err := property.ParseType(cpr.Type); err != nil {
		e.Add("type", fmt.Sprintf("invalid property type [%s, %s]", property.TypeHome, property.TypeApartment))
	}

	if property.Type(cpr.Type) == property.TypeApartment {
		if cpr.UnitNumber == nil {
			e.Add("unit number", "unit number required for apartments")
		}
	}

	return e
}

// CreatePropertyResponse - represents a client create property response.
type CreatePropertyResponse struct {
	Property ClientProperty `json:"property"`
}

// CreateProperty - invoked by POST /v1/managers/:id/properties.
func (c *Controller) CreateProperty(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	c.log.Info().Msg("Creating Property.")

	var input CreatePropertyRequest
	if err := api.Decode(r, &input); err != nil {
		c.log.Err(err).Msg("Unable to decode create property request.")
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

	p, err := c.Property.Create(ctx, toCoreNewProperty(input), mID)
	if err != nil {
		c.log.Err(err).Msg("Unable to create property.")
		return api.InternalServerError("Error.", err, nil)
	}

	c.log.Info().Msg("Successfully created Property.")
	return api.Respond(w, http.StatusCreated, CreatePropertyResponse{Property: toClientProperty(p)})
}

// UpdatePropertyRequest - represents input for updating an existing property.
type UpdatePropertyRequest struct {
	Name *string  `json:"name"`
	Rent *float64 `json:"rent"`
}

func (upr *UpdatePropertyRequest) validate() *api.ValidationError {
	e := api.NewValidationError()
	if upr.Name != nil && *upr.Name == "" {
		e.Add("name", "cannot be empty")
	}
	if upr.Rent != nil && *upr.Rent <= 0 {
		e.Add("rent", "needs to be greater than 0")
	}
	if upr.Name == nil && upr.Rent == nil {
		e.Add("input", "must provided update values")
	}
	return e
}

// UpdatePropertyResponse - represents a client update property response.
type UpdatePropertyResponse struct {
	Property ClientProperty `json:"property"`
}

// UpdateProperty - invoked by PATCH /v1/managers/:id/properties/:id.
func (c *Controller) UpdateProperty(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	c.log.Info().Msg("Updating Property.")

	var input UpdatePropertyRequest
	if err := api.Decode(r, &input); err != nil {
		c.log.Err(err).Msg("Unable to decode update property request.")
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

	p, err := c.Property.Update(ctx, pID, toCoreUpdateProperty(input))
	if err != nil {
		c.log.Err(err).Msg("Unable to update property.")
		if errors.Is(err, core.ErrNotFound) {
			return api.NotFoundError("Item not found.", err, nil)
		}
		return api.InternalServerError("Error.", err, nil)
	}

	c.log.Info().Msg("Successfully updated Property.")
	return api.Respond(w, http.StatusOK, UpdatePropertyResponse{Property: toClientProperty(p)})
}

// GetPropertyResponse - represents a client get property response.
type GetPropertyResponse struct {
	Property ClientProperty `json:"property"`
}

// GetProperty - invoked by GET /v1/managers/:id/properties/:id.
func (c *Controller) GetProperty(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	c.log.Info().Msg("Fetching Property.")

	if _, err := uuid.Parse(api.PathParam(r, managerIDPathParam)); err != nil {
		c.log.Err(err).Msg("Invalid manager id.")
		return api.BadRequestError("Invalid id.", err, nil)
	}

	pID, err := uuid.Parse(api.PathParam(r, propertyIDPathParam))
	if err != nil {
		c.log.Err(err).Msg("Invalid property id.")
		return api.BadRequestError("Invalid id.", err, nil)
	}

	p, err := c.Property.Get(ctx, pID)
	if err != nil {
		c.log.Err(err).Msg("Unable to get property.")
		if errors.Is(err, core.ErrNotFound) {
			return api.NotFoundError("Item not found.", err, nil)
		}
		return api.InternalServerError("Error.", err, nil)
	}

	c.log.Info().Msg("Successfully Fetched Property.")
	return api.Respond(w, http.StatusOK, CreatePropertyResponse{Property: toClientProperty(p)})
}

// ListPropertiesResponse - represents a client list property response.
type ListPropertiesResponse struct {
	Properties []ClientProperty `json:"properties"`
}

// ListProperties - invoked by GET /v1/managers/:id/properties.
func (c *Controller) ListProperties(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	c.log.Info().Msg("Listing Properties.")

	mID, err := uuid.Parse(api.PathParam(r, managerIDPathParam))
	if err != nil {
		c.log.Err(err).Msg("Invalid manager id.")
		return api.BadRequestError("Invalid id.", err, nil)
	}

	pps, err := c.Property.List(ctx, mID)
	if err != nil {
		c.log.Err(err).Msg("Unable to list properties.")
		return api.InternalServerError("Error.", err, nil)
	}

	c.log.Info().Msg("Successfully Listed Properties.")
	return api.Respond(w, http.StatusOK, ListPropertiesResponse{Properties: toClientProperties(pps)})
}
