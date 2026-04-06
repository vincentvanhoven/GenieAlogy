-- +goose Up
CREATE TABLE lamp_meta
(
    id TEXT PRIMARY KEY,
    value TEXT
);

-- +goose Down
DROP TABLE lamp_meta;
