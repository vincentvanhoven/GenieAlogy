-- +goose Up

UPDATE people SET `parent_arrow_position` = 'top';

-- +goose Down
UPDATE people SET `parent_arrow_position` = null;
