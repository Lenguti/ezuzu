package v1

import (
	"github.com/google/uuid"
	"github.com/lenguti/ezuzu/business/core/tenant"
)

// ClientTenant - represents a client tenant entity.
type ClientTenant struct {
	ID          string `json:"id"`
	Type        string `json:"type"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	DateOfBirth string `json:"dateOfBirth"`
	SSN         string `json:"ssn"`
	CreatedAt   int64  `json:"createdAt"`
	UpdatedAt   int64  `json:"updatedAt"`
}

func toCoreNewTenant(input CreateTenantRequest) tenant.NewTenant {
	newTenant := tenant.NewTenant{
		FirstName:   input.FirstName,
		LastName:    input.LastName,
		DateOfBirth: input.DateOfBirth,
		SSN:         input.SSN,
		Type:        tenant.Type(input.Type),
	}
	return newTenant
}

func toCoreUpdateTenant(input UpdateTenantRequest) tenant.UpdateTenant {
	var updateTenant tenant.UpdateTenant
	if input.NewPropertyID != nil {
		uid := uuid.MustParse(*input.NewPropertyID)
		updateTenant.PropertyID = &uid
	}
	if input.Type != nil {
		updateTenant.Type = tenant.ToPtrType(*input.Type)
	}
	return updateTenant
}

func toClientTenants(input []tenant.Tenant) []ClientTenant {
	ctts := make([]ClientTenant, 0, len(input))
	for _, ct := range input {
		ctts = append(ctts, toClientTenant(ct))
	}
	return ctts
}

func toClientTenant(input tenant.Tenant) ClientTenant {
	return ClientTenant{
		ID:          input.ID.String(),
		Type:        input.Type.String(),
		FirstName:   input.FirstName,
		LastName:    input.LastName,
		DateOfBirth: input.DateOfBirth,
		SSN:         "#########",
		CreatedAt:   input.CreatedAt.Unix(),
		UpdatedAt:   input.UpdatedAt.Unix(),
	}
}
