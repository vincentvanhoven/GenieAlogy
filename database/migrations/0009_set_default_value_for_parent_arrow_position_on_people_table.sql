-- +goose Up

ALTER TABLE people RENAME COLUMN `parent_arrow_position` to `parent_arrow_position_old`;
ALTER TABLE people ADD COLUMN `parent_arrow_position` DEFAULT 'top';
UPDATE people SET `parent_arrow_position` = `parent_arrow_position_old`;
ALTER TABLE people DROP COLUMN `parent_arrow_position_old`;

-- +goose Down
ALTER TABLE people RENAME COLUMN `parent_arrow_position` to `parent_arrow_position_old`;
ALTER TABLE people ADD COLUMN `parent_arrow_position`;
UPDATE people SET `parent_arrow_position` = `parent_arrow_position_old`;
ALTER TABLE people DROP COLUMN `parent_arrow_position_old`;
