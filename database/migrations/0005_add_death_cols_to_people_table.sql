-- +goose Up
ALTER TABLE people ADD COLUMN deathdate TEXT;
ALTER TABLE people ADD COLUMN deathplace TEXT;

-- +goose Down
ALTER TABLE people DROP COLUMN deathdate;
ALTER TABLE people DROP COLUMN deathplace;
