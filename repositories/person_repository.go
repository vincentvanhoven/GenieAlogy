package repositories

import (
	"GenieAlogy/models"
	"database/sql"
	"errors"
	"fmt"
	"math"
)

type PersonRepository struct {
}

var PersonRepo = &PersonRepository{}

func (repo *PersonRepository) Create(p models.Person) (*int, error) {
	var lastInsertId int64

	// Open transaction
	transaction, err := DatabaseRepo.DB.Begin()
	if err != nil {
		return nil, err
	}

	var result sql.Result

	// Attempt the execution of the prepared statement
	result, err = transaction.Exec(
		`INSERT INTO people (
				sex, firstname, lastname, birthdate, birthplace, family_id, position_x, position_y
			)
         	VALUES (?, ?, ?, ?, ?, ?, ?, ?)`,
		p.Sex, p.Firstname, p.Lastname, p.Birthdate, p.Birthplace, p.FamilyId, p.PositionX, p.PositionY,
	)

	// Rollback if anything went wrong
	if err != nil {
		rollbackerr := transaction.Rollback()
		if rollbackerr != nil {
			return nil, rollbackerr
		}

		return nil, err
	}

	// All went well, commit the transaction (this may return an error)
	err = transaction.Commit()

	// Handle the error, if anything went wrong
	if err != nil {
		return nil, err
	}

	// Attempt to get the Id that was inserted
	lastInsertId, err = result.LastInsertId()

	// Handle the error, if anything went wrong
	if err != nil {
		return nil, err
	}

	// If lastInsertId cannot be cast to int64, thrown an error
	if lastInsertId < math.MinInt || lastInsertId > math.MaxInt {
		return nil, errors.New(fmt.Sprintf("Last insert id out of range: %d", lastInsertId))
	}

	// Else, cast and retun it
	var castInsertId = int(lastInsertId)
	return &castInsertId, nil
}

func (repo *PersonRepository) Fetch(id int) (*models.Person, error) {
	var p models.Person

	err := DatabaseRepo.DB.
		QueryRow(`SELECT * FROM people WHERE id = ?`, id).
		Scan(&p.Id, &p.Sex, &p.Firstname, &p.Lastname, &p.Birthdate, &p.Birthplace, &p.FamilyId, &p.ProfilePicture, &p.PositionX, &p.PositionY)

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
			&row.Id,
			&row.Sex,
			&row.Firstname,
			&row.Lastname,
			&row.Birthdate,
			&row.Birthplace,
			&row.FamilyId,
			&row.ProfilePicture,
			&row.PositionX,
			&row.PositionY,
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
			SET sex=?, firstname=?, lastname=?, birthdate=?, birthplace=?, family_id=?, profile_picture=?, position_x=?, position_y=?
			WHERE id=?
		`,
		p.Sex, p.Firstname, p.Lastname, p.Birthdate, p.Birthplace, p.FamilyId, p.ProfilePicture, p.PositionX, p.PositionY, p.Id,
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

func (repo *PersonRepository) Delete(id int) error {
	_, err := DatabaseRepo.DB.Exec(
		`DELETE FROM people WHERE id=?`,
		id,
	)

	return err
}
