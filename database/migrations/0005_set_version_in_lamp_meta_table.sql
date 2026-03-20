-- +goose Up

INSERT INTO lamp_meta (id, value) VALUES ('version', '0.0.1');

-- +goose Down
DELETE FROM lamp_meta WHERE id = 'version';
