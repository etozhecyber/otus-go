-- +goose Up
ALTER TABLE IF EXISTS events
ADD COLUMN IF NOT EXISTS is_notified BOOLEAN NOT NULL DEFAULT 'false';

-- +goose Down
ALTER TABLE IF EXISTS events
DROP COLUMN IF EXISTS is_notified;
