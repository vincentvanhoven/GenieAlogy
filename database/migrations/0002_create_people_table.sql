-- +goose Up
CREATE TABLE people
(
    uuid            TEXT PRIMARY KEY,
    sex             TEXT CHECK (sex IN ('male', 'female')),
    firstname       TEXT,
    lastname        TEXT,
    birthdate       TEXT,
    birthplace      TEXT,
    family_uuid     TEXT,
    profile_picture BLOB,
    position_x      INTEGER,
    position_y      INTEGER
--     FOREIGN KEY (family_uuid) REFERENCES families (uuid) ON DELETE RESTRICT
);

-- +goose Down
DROP TABLE people;
