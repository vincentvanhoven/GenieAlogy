-- +goose Up

UPDATE lamp_meta SET `value` = '0.0.2' WHERE `id` = 'version';

-- +goose Down
UPDATE lamp_meta SET `value` = '0.0.1' WHERE `id` = 'version';
