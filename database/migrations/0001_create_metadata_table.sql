-- +goose Up
CREATE TABLE metadata
(
    id    INTEGER PRIMARY KEY AUTOINCREMENT,
    value TEXT
);

-- +goose Down
DROP TABLE metadata;
