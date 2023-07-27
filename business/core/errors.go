package core

// Error represents custom business logic errors.
type Error string

// Error satisfies the error interface.
func (e Error) Error() string {
	return string(e)
}

const (
	// ErrNotFound represents an item not found.
	ErrNotFound = Error("item not found")
)
