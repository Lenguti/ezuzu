-- +goose Up
-- +goose StatementBegin
CREATE TABLE tenant (
    id uuid NOT NULL,
    property_id uuid NOT NULL,
    type text,
    first_name text,
    last_name text,
    date_of_birth text,
    ssn int,
    created_at int,
    updated_at int,
    PRIMARY KEY (id),
    FOREIGN KEY(property_id) REFERENCES property(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE tenant;
-- +goose StatementEnd
