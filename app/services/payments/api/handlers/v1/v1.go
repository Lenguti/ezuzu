package v1

import (
	"context"
	"fmt"
	"net/http"

	"github.com/lenguti/ezuzu/foundation/api"
)

const (
	managerIDPathParam  = "managerId"
	propertyIDPathParam = "propertyId"
	tenantIDPathParam   = "tenantId"
	invoiceIDPathParam  = "invoiceId"
)

// Routes - route definitions for v1.
func (c *Controller) Routes() *api.Router {
	const version = "v1"

	c.router.Handle(http.MethodGet, version, "/status", c.status)

	c.router.Handle(http.MethodPost, version, "/managers/:managerId/properties/:propertyId/invoices", c.CreateInvoice)
	c.router.Handle(http.MethodPost, version, "/tenants/:tenantId/invoices/:invoiceId/payments", c.CreatePayment)
	c.router.Handle(http.MethodGet, version, "/tenants/:tenantId/payments/history", c.ListPaymentHistory)

	return c.router
}

func (c *Controller) status(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	if err := c.db.Connect(); err != nil {
		c.log.Error().Err(err).Msg("Unable to connect to db.")
		return fmt.Errorf("status: unable to connect to db: %w", err)
	}
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(`{"status": "ok"}`))
	return nil
}
