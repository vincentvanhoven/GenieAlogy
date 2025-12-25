package repositories

import (
	"GenieAlogy/models"
)

type PersonRepository struct {
}

var PersonRepo = &PersonRepository{}

func (repo *PersonRepository) Create(p models.Person) error {
	// Open transaction
	transaction, err := DatabaseRepo.DB.Begin()
	if err != nil {
		return err
	}

	// Attempt the execution of the prepared statement
	_, err = transaction.Exec(
		`INSERT INTO people (
				uuid, sex, firstname, lastname, birthdate, birthplace, family_uuid, position_x, position_y
			)
         	VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		p.Uuid, p.Sex, p.Firstname, p.Lastname, p.Birthdate, p.Birthplace, p.FamilyUuid, p.Position.X, p.Position.Y,
	)

	// Rollback if anything went wrong
	if err != nil {
		rollbackerr := transaction.Rollback()
		if rollbackerr != nil {
			return rollbackerr
		}

		return err
	}

	// All went well, commit the transaction (this may return an error)
	return transaction.Commit()
}

func (repo *PersonRepository) Fetch(uuid string) (*models.Person, error) {
	var p models.Person

	err := DatabaseRepo.DB.
		QueryRow(`SELECT * FROM people WHERE uuid = ?`, uuid).
		Scan(&p.Uuid, &p.Sex, &p.Firstname, &p.Lastname, &p.Birthdate, &p.Birthplace, &p.FamilyUuid, &p.ProfilePicture, &p.Position.X, &p.Position.Y)

	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (repo *PersonRepository) FetchAll() ([]models.Person, error) {
	var p []models.Person

	rows, err := DatabaseRepo.DB.Query(`SELECT * FROM people`)
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

func (repo *PersonRepository) Update(p models.Person) error {
	// Open transaction
	transaction, err := DatabaseRepo.DB.Begin()
	if err != nil {
		return err
	}

	// Attempt the execution of the prepared statement
	_, err = transaction.Exec(
		`
			UPDATE people
			SET sex=?, firstname=?, lastname=?, birthdate=?, birthplace=?, family_uuid=?, profile_picture=?, position_x=?, position_y=?
			WHERE uuid=?
		`,
		p.Sex, p.Firstname, p.Lastname, p.Birthdate, p.Birthplace, p.FamilyUuid, p.ProfilePicture, p.Position.X, p.Position.Y, p.Uuid, p.Uuid,
	)

	// Rollback if anything went wrong
	if err != nil {
		rollbackerr := transaction.Rollback()
		if rollbackerr != nil {
			return rollbackerr
		}

		return err
	}

	// All went well, commit the transaction (this may return an error)
	return transaction.Commit()
}

func (repo *PersonRepository) Delete(uuid string) error {
	_, err := DatabaseRepo.DB.Exec(
		`DELETE FROM people WHERE uuid=?`,
		uuid,
	)

	return err
}
