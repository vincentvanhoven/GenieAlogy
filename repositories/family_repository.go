package repositories

import (
	"GenieAlogy/models"
	"database/sql"
	"errors"
	"fmt"
	"math"
)

type FamilyRepository struct{}

var FamilyRepo = &FamilyRepository{}

func (repo *FamilyRepository) Create(f models.Family) (*int, error) {
	var lastInsertId int64

	// Open transaction
	transaction, err := DatabaseRepo.DB.Begin()
	if err != nil {
		return nil, err
	}

	var result sql.Result

	// Attempt the execution of the prepared statement
	result, err = transaction.Exec(
		`INSERT INTO families (person_1_id, person_2_id, position_x, position_y) VALUES (?, ?, ?, ?)`,
		f.Person1Id, f.Person2Id, f.PositionX, f.PositionY,
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

func (repo *FamilyRepository) Fetch(id int) (*models.Family, error) {
	var f models.Family

	err := DatabaseRepo.DB.QueryRow(
		`SELECT
    		id, person_1_id, person_2_id, position_x, position_y
         	FROM families
         	WHERE id = ?
		`,
		id,
	).Scan(&f.Id, &f.Person1Id, &f.Person2Id, &f.PositionX, &f.PositionY)

	if err != nil {
		return nil, err
	}
	return &f, nil
}

func (repo *FamilyRepository) FetchAll() ([]models.Family, error) {
	var f []models.Family

	rows, err := DatabaseRepo.DB.Query(`SELECT * FROM families`)
	if err != nil {
		return f, err
	}

	defer rows.Close()

	for rows.Next() {
		var row models.Family
		err := rows.Scan(&row.Id, &row.Person1Id, &row.Person2Id, &row.PositionX, &row.PositionY)

		if err != nil {
			return f, err
		}
		f = append(f, row)
	}

	// check for errors after iteration
	if err = rows.Err(); err != nil {
		return f, err
	}

	return f, nil
}

func (repo *FamilyRepository) Update(f models.Family) error {
	// Open transaction
	transaction, err := DatabaseRepo.DB.Begin()
	if err != nil {
		return err
	}

	// Attempt the execution of the prepared statement
	_, err = transaction.Exec(
		`UPDATE families SET person_1_id = ?, person_2_id = ?, position_x = ?, position_y = ? WHERE id = ?`,
		f.Person1Id, f.Person2Id, f.PositionX, f.PositionY, f.Id,
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

func (repo *FamilyRepository) Delete(id int) error {
	transaction, err := DatabaseRepo.DB.Begin()
	if err != nil {
		return err
	}

	_, err = transaction.Exec(
		`DELETE FROM families WHERE id=?`,
		id,
	)

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
