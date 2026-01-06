-- +goose Up
CREATE TABLE metadata
(
    id    TEXT PRIMARY KEY,
    value TEXT
);

-- +goose Down
DROP TABLE metadata;
