-- +goose Up
-- +goose StatementBegin
CREATE TABLE payments (
    id uuid NOT NULL,
    invoice_id uuid NOT NULL,
    tenant_id uuid NOT NULL,
    amount float,
    created_at int,
    updated_at int,
    PRIMARY KEY (id),
    FOREIGN KEY(invoice_id) REFERENCES invoices(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE payments;
-- +goose StatementEnd
