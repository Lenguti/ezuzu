package v1

import "github.com/lenguti/ezuzu/business/core/property"

// ClientProperty - represents a client property entity.
type ClientProperty struct {
	ID         string  `json:"id"`
	Address    string  `json:"address"`
	Name       string  `json:"name"`
	Rent       float64 `json:"rent"`
	Type       string  `json:"type"`
	UnitNumber *int    `json:"unitNumber,omitempty"`
	CreatedAt  int64   `json:"createdAt"`
	UpdatedAt  int64   `json:"updatedAt"`
}

func toCoreNewProperty(input CreatePropertyRequest) property.NewProperty {
	newProperty := property.NewProperty{
		Address:    input.formattedAddress,
		Name:       input.Name,
		Rent:       input.Rent,
		Type:       property.Type(input.Type),
		UnitNumber: input.UnitNumber,
	}
	return newProperty
}

func toCoreUpdateProperty(input UpdatePropertyRequest) property.UpdateProperty {
	updateProperty := property.UpdateProperty{
		Name: input.Name,
		Rent: input.Rent,
	}
	return updateProperty
}

func toClientProperties(input []property.Property) []ClientProperty {
	cpps := make([]ClientProperty, 0, len(input))
	for _, pp := range input {
		cpps = append(cpps, toClientProperty(pp))
	}
	return cpps
}

func toClientProperty(input property.Property) ClientProperty {
	return ClientProperty{
		ID:         input.ID.String(),
		Address:    input.Address,
		Name:       input.Name,
		Rent:       input.Rent,
		Type:       input.Type.String(),
		UnitNumber: input.UnitNumber,
		CreatedAt:  input.CreatedAt.Unix(),
		UpdatedAt:  input.UpdatedAt.Unix(),
	}
}
