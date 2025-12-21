package repositories

import (
	"GenieAlogy/models"
	"database/sql"
)

type PersonRepository struct {
	DB *sql.DB
}

func (r *PersonRepository) Insert(p models.Person) error {
	_, err := r.DB.Exec(
		`
			INSERT INTO people (
				uuid, sex, firstname, lastname, birthdate, birthplace, family_uuid, position_x, position_y
			)
         	VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
		`,
		p.Uuid, p.Sex, p.Firstname, p.Lastname, p.Birthdate, p.Birthplace, p.FamilyUuid, p.Position.X, p.Position.Y,
	)

	return err
}

func (r *PersonRepository) Fetch(uuid string) (*models.Person, error) {
	var p models.Person

	err := r.DB.
		QueryRow(`SELECT * FROM people WHERE uuid = ?`, uuid).
		Scan(&p.Uuid, &p.Sex, &p.Firstname, &p.Lastname, &p.Birthdate, &p.Birthplace, &p.FamilyUuid, &p.ProfilePicture, &p.Position.X, &p.Position.Y)

	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *PersonRepository) FetchAll() ([]models.Person, error) {
	var p []models.Person

	rows, err := r.DB.Query(`SELECT * FROM people`)
	if err != nil {
		return p, err
	}

	defer rows.Close()

	for rows.Next() {
		var row models.Person
		err := rows.Scan(
			&row.Uuid,
			&row.Sex,
			&row.Firstname,
			&row.Lastname,
			&row.Birthdate,
			&row.Birthplace,
			&row.FamilyUuid,
			&row.ProfilePicture,
			&row.Position.X,
			&row.Position.Y,
		)

		if err != nil {
			return p, err
		}
		p = append(p, row)
	}

	// check for errors after iteration
	if err = rows.Err(); err != nil {
		return p, err
	}

	return p, nil
}

func (r *PersonRepository) Update(p models.Person) error {
	_, err := r.DB.Exec(
		`
			UPDATE people
			SET sex=?, firstname=?, lastname=?, birthdate=?, birthplace=?, family_uuid=?, profile_picutre=?, position_x=?, position_y=?
			WHERE uuid=?
		`,
		p.Sex, p.Firstname, p.Lastname, p.Birthdate, p.Birthplace, p.FamilyUuid, p.ProfilePicture, p.Position.X, p.Position.Y, p.Uuid, p.Uuid,
	)

	return err
}

func (r *PersonRepository) Delete(uuid string) error {
	_, err := r.DB.Exec(
		`DELETE FROM people WHERE uuid=?`,
		uuid,
	)

	return err
}
