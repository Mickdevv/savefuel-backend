
-- +goose Up
ALTER TABLE users ADD COLUMN login_attempts int not null default 0;

-- +goose Down
ALTER TABLE users DROP COLUMN login_attempts;

