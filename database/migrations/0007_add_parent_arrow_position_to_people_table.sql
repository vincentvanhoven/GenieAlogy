-- +goose Up

ALTER TABLE people ADD COLUMN parent_arrow_position TEXT;

-- +goose Down
ALTER TABLE people DROP COLUMN parent_arrow_position;
