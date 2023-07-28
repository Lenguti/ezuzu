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

	// ErrPastDuePayment represents an error for payment past due.
	ErrPastDuePayment = Error("payment cannot be past due")

	// ErrPaymentConflict represents an error for payment amount mismatch.
	ErrPaymentConflict = Error("payments must not exceed invoice amount")
)
