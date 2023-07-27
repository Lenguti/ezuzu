package property

import (
	"fmt"
)

// Type - represents property type enum.
type Type string

// String - returns string representation of type.
func (t Type) String() string {
	return string(t)
}

const (
	TypeHome      = "HOME"
	TypeApartment = "APARTMENT"
)

var validPropertyType = map[Type]struct{}{
	TypeHome:      {},
	TypeApartment: {},
}

// ParseType - will attempt to validate the provided type.
func ParseType(v string) error {
	if _, ok := validPropertyType[Type(v)]; !ok {
		return fmt.Errorf("parse type: invalid property type")
	}
	return nil
}
