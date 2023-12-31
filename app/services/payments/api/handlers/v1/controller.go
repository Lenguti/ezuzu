package v1

import (
	"fmt"

	"github.com/lenguti/ezuzu/app/services/property/api/handlers/v1/client"
	"github.com/lenguti/ezuzu/business/core/invoice"
	"github.com/lenguti/ezuzu/business/core/invoice/stores/invoicedb"
	"github.com/lenguti/ezuzu/business/core/payment"
	"github.com/lenguti/ezuzu/business/core/payment/stores/paymentdb"
	paymenthistory "github.com/lenguti/ezuzu/business/core/payment_history"
	paymenthistorydb "github.com/lenguti/ezuzu/business/core/payment_history/stores/payment_historydb"
	"github.com/lenguti/ezuzu/business/data/db"
	"github.com/lenguti/ezuzu/foundation/api"
	"github.com/rs/zerolog"
)

// Controller - represents our handler service orchestrator.
type Controller struct {
	Invoice        *invoice.Core
	Payment        *payment.Core
	PaymentHistory *paymenthistory.Core
	PC             client.IProperty

	db     *db.DB
	config Config
	log    zerolog.Logger
	router *api.Router
}

// NewController - initializes a new controller with all its services.
func NewController(log zerolog.Logger, cfg Config, pc client.IProperty) (*Controller, error) {
	ddb, err := db.New(db.Config{
		User:         cfg.DBUser,
		Password:     cfg.DBPass,
		Name:         cfg.DBName,
		Host:         cfg.DBHost,
		Port:         cfg.DBPort,
		MaxIdleConns: 10,
		MaxOpenConns: 10,
	})
	if err != nil {
		return nil, fmt.Errorf("new controller: unable to initialize new db: %w", err)
	}

	ic := invoice.NewCore(invoicedb.NewStore(ddb), log)
	pymc := payment.NewCore(paymentdb.NewStore(ddb), ic, log)
	phc := paymenthistory.NewCore(paymenthistorydb.NewStore(ddb), log)

	return &Controller{db: ddb,
		Invoice:        ic,
		Payment:        pymc,
		PaymentHistory: phc,
		PC:             pc,

		config: cfg,
		log:    log,
		router: api.NewRouter(),
	}, nil
}
