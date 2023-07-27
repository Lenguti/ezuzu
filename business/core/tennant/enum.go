package tennant

import "fmt"

// Type - represents tennant type enum.
type Type string

// String - returns string representation of type.
func (t Type) String() string {
	return string(t)
}

const (
	TypePrimary   = "PRIMARY"
	TypeSecondary = "SECONDARY"
)

var validTennantType = map[Type]struct{}{
	TypePrimary:   {},
	TypeSecondary: {},
}

// ParseType - will attempt to validate the provided type.
func ParseType(v string) error {
	if _, ok := validTennantType[Type(v)]; !ok {
		return fmt.Errorf("parse type: invalid tennant type")
	}
	return nil
}
