-- +goose Up
CREATE TABLE families
(
    uuid          TEXT PRIMARY KEY,
    person_1_uuid TEXT,
    person_2_uuid TEXT,
    position_x    INTEGER,
    position_y    INTEGER
--     FOREIGN KEY (person_1_uuid) REFERENCES people (uuid) ON DELETE RESTRICT,
--     FOREIGN KEY (person_2_uuid) REFERENCES people (uuid) ON DELETE RESTRICT
);

-- +goose Down
DROP TABLE families;
