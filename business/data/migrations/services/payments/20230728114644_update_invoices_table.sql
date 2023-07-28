-- +goose Up
-- +goose StatementBegin
ALTER TABLE invoices
ADD manager_id uuid NOT NULL; 
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE invoices
DROP COLUMN manager_id; 
-- +goose StatementEnd
