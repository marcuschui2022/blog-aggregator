-- +goose Up
CREATE TABLE users
(
    id         uuid,
    created_at timestamp,
    update_at  timestamp,
    name       text unique
);

-- +goose Down
DROP TABLE users;