-- +goose Up
-- +goose StatementBegin
ALTER TABLE property
ADD rent float; 
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE property
DROP COLUMN rent; 
-- +goose StatementEnd
