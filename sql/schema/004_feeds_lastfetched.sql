-- +goose Up
alter table feeds
    add column last_fetched_at timestamp;

-- +goose Down
ALTER TABLE feeds DROP COLUMN last_fetched_at;