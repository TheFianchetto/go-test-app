-- +goose Up
-- +goose StatementBegin
ALTER TABLE users ADD COLUMN username VARCHAR(255) NOT NULL DEFAULT 'Test';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE users DROP COLUMN username;
-- +goose StatementEnd
