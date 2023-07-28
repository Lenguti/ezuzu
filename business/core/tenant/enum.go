package tenant

import "fmt"

// Type - represents tenant type enum.
type Type string

// String - returns string representation of type.
func (t Type) String() string {
	return string(t)
}

func ToPtrType(v string) *Type {
	t := Type(v)
	return &t
}

const (
	TypePrimary   = "PRIMARY"
	TypeSecondary = "SECONDARY"
)

var validTenantType = map[Type]struct{}{
	TypePrimary:   {},
	TypeSecondary: {},
}

// ParseType - will attempt to validate the provided type.
func ParseType(v string) error {
	if _, ok := validTenantType[Type(v)]; !ok {
		return fmt.Errorf("parse type: invalid tenant type")
	}
	return nil
}
