-- +goose Up
-- +goose StatementBegin
CREATE TABLE property (
    id uuid NOT NULL,
    manager_id uuid NOT NULL,
    address text,
    name text,
    type text,
    unit_number int NULL,
    created_at int,
    updated_at int,
    PRIMARY KEY (id),
    FOREIGN KEY(manager_id) REFERENCES property_manager(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE property;
-- +goose StatementEnd
