package paymenthistory

import "time"

// PaymentHistory - represents a business domain payment history.
type PaymentHistory struct {
	Month     string
	Total     float64
	Timestamp time.Time
}
