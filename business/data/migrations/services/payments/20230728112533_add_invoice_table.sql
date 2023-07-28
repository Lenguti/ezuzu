-- +goose Up
-- +goose StatementBegin
CREATE TABLE invoices (
    id uuid NOT NULL,
    property_id uuid NOT NULL,
    tenant_id uuid NOT NULL,
    amount float,
    due_date int,
    created_at int,
    updated_at int,
    PRIMARY KEY (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE invoices;
-- +goose StatementEnd
