package repositories

import (
	"GenieAlogy/models"
)

type FamilyRepository struct{}

var FamilyRepo = &FamilyRepository{}

func (repo *FamilyRepository) Create(f models.Family) error {
	// Open transaction
	transaction, err := DatabaseRepo.DB.Begin()
	if err != nil {
		return err
	}

	// Attempt the execution of the prepared statement
	_, err = transaction.Exec(
		`INSERT INTO families (uuid, person_1_uuid, person_2_uuid, position_x, position_y) VALUES (?, ?, ?, ?, ?)`,
		f.Uuid, f.Person1Uuid, f.Person2Uuid, f.PositionX, f.PositionY,
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

func (repo *FamilyRepository) Fetch(uuid string) (*models.Family, error) {
	var f models.Family

	err := DatabaseRepo.DB.QueryRow(
		`SELECT
    		uuid, person_1_uuid, person_2_uuid, position_x, position_y
         	FROM families
         	WHERE uuid = ?
		`,
		uuid,
	).Scan(&f.Uuid, &f.Person1Uuid, &f.Person2Uuid, &f.PositionX, &f.PositionY)

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
		err := rows.Scan(&row.Uuid, &row.Person1Uuid, &row.Person2Uuid, &row.PositionX, &row.PositionY)

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
		`UPDATE families SET person_1_uuid = ?, person_2_uuid = ?, position_x = ?, position_y = ? WHERE uuid = ?`,
		f.Person1Uuid, f.Person2Uuid, f.PositionX, f.PositionY, f.Uuid,
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

func (repo *FamilyRepository) Delete(uuid string) error {
	transaction, err := DatabaseRepo.DB.Begin()
	if err != nil {
		return err
	}

	_, err = transaction.Exec(
		`DELETE FROM families WHERE uuid=?`,
		uuid,
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
