-- +goose Up
-- +goose StatementBegin
CREATE OR REPLACE VIEW payment_history AS
SELECT
    tenant_id AS id,
	SUM(amount) AS total,
    DATE_TRUNC('month', TO_TIMESTAMP(created_at)) as month,
    MIN(created_at) AS earliest_payment
FROM
    payments
GROUP BY
    tenant_id,
    month
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP VIEW payment_history;
-- +goose StatementEnd
