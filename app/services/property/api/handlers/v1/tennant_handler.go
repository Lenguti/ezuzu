package v1

import (
	"context"
	"fmt"
	"net/http"
	"regexp"
	"strconv"

	"github.com/google/uuid"
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
