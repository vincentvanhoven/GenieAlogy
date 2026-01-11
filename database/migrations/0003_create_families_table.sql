-- +goose Up
CREATE TABLE families
(
    id          INTEGER PRIMARY KEY AUTOINCREMENT,
    person_1_id INTEGER,
    person_2_id INTEGER,
    position_x  INTEGER,
    position_y  INTEGER
--     FOREIGN KEY (person_1_id) REFERENCES people (id) ON DELETE RESTRICT,
--     FOREIGN KEY (person_2_id) REFERENCES people (id) ON DELETE RESTRICT
);

-- +goose Down
DROP TABLE families;
