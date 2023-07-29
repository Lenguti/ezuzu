package v1

import (
	paymenthistory "github.com/lenguti/ezuzu/business/core/payment_history"
)

// ClientPaymentHistory - represents a client payment history entity.
type ClientPaymentHistory struct {
	Month     string  `json:"month"`
	Total     float64 `json:"total"`
	Timestamp int64   `json:"timestamp"`
}

func toClientPaymentHistories(phs []paymenthistory.PaymentHistory) []ClientPaymentHistory {
	cphs := make([]ClientPaymentHistory, 0, len(phs))
	for _, ph := range phs {
		cphs = append(cphs, toClientPaymentHistory(ph))
	}
	return cphs
}

func toClientPaymentHistory(ph paymenthistory.PaymentHistory) ClientPaymentHistory {
	return ClientPaymentHistory{
		Month:     ph.Month,
		Total:     ph.Total,
		Timestamp: ph.Timestamp.Unix(),
	}
}
