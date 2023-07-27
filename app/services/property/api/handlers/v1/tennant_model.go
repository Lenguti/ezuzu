package v1

import "github.com/lenguti/ezuzu/business/core/tennant"

// ClientTennant - represents a client tennant entity.
type ClientTennant struct {
	ID          string `json:"id"`
	Type        string `json:"type"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	DateOfBirth string `json:"dateOfBirth"`
	SSN         string `json:"ssn"`
	CreatedAt   int64  `json:"createdAt"`
	UpdatedAt   int64  `json:"updatedAt"`
}

func toCoreNewTennant(input CreateTennantRequest) tennant.NewTennant {
	newTennant := tennant.NewTennant{
		FirstName:   input.FirstName,
		LastName:    input.LastName,
		DateOfBirth: input.DateOfBirth,
		SSN:         input.SSN,
		Type:        tennant.Type(input.Type),
	}
	return newTennant
}

func toClientTennants(input []tennant.Tennant) []ClientTennant {
	ctts := make([]ClientTennant, 0, len(input))
	for _, ct := range input {
		ctts = append(ctts, toClientTennant(ct))
	}
	return ctts
}

func toClientTennant(input tennant.Tennant) ClientTennant {
	return ClientTennant{
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
