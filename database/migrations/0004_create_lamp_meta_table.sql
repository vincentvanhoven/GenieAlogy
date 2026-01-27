-- +goose Up
CREATE TABLE lamp_meta
(
    id TEXT PRIMARY KEY,
    value TEXT
);

INSERT INTO lamp_meta (id, value) VALUES ('version', '0.0.1');

-- +goose Down
DROP TABLE lamp_meta;
