-- +goose Up
CREATE TABLE people
(
    id              INTEGER PRIMARY KEY AUTOINCREMENT,
    sex             TEXT CHECK (sex IN ('male', 'female')),
    firstname       TEXT,
    lastname        TEXT,
    birthdate       TEXT,
    birthplace      TEXT,
    family_id       INTEGER,
    profile_picture BLOB,
    position_x      INTEGER,
    position_y      INTEGER
--     FOREIGN KEY (family_id) REFERENCES families (id) ON DELETE RESTRICT
);

-- +goose Down
DROP TABLE people;
