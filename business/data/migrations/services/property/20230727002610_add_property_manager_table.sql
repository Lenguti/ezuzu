-- +goose Up
-- +goose StatementBegin
CREATE TABLE property_manager (
    id uuid NOT NULL,
    entity text,
    created_at int,
    updated_at int,
    PRIMARY KEY (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE property_manager;
-- +goose StatementEnd
