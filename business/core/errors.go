package core

// Error represents custom business logic errors.
type Error string

// Error satisfies the error interface.
func (e Error) Error() string {
	return string(e)
}
