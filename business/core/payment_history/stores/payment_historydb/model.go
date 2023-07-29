package paymenthistorydb

import (
	"fmt"
	"time"

	paymenthistory "github.com/lenguti/ezuzu/business/core/payment_history"
)

type dbPaymentHistory struct {
	ID        string    `db:"id"`
	Month     time.Time `db:"month"`
	Total     float64   `db:"total"`
	Timestamp int64     `db:"earliest_payment"`
}

func toCorePaymentHistories(dbphs []dbPaymentHistory) []paymenthistory.PaymentHistory {
	phs := make([]paymenthistory.PaymentHistory, 0, len(dbphs))
	for _, dbph := range dbphs {
		phs = append(phs, toCorePaymentHistory(dbph))
	}
	return phs
}

func toCorePaymentHistory(dbph dbPaymentHistory) paymenthistory.PaymentHistory {
	return paymenthistory.PaymentHistory{
		Month:     fmt.Sprintf("%v - %v", dbph.Month.Month(), dbph.Month.Year()),
		Total:     dbph.Total,
		Timestamp: time.Unix(dbph.Timestamp, 0),
	}
}
