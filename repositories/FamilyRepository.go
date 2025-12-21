package repositories

import (
	"GenieAlogy/models"
	"database/sql"
)

type FamilyRepository struct {
	DB *sql.DB
}

func (r *FamilyRepository) Insert(f models.Family) error {
	_, err := r.DB.Exec(
		`
			INSERT INTO families (uuid, person_1_uuid, person_2_uuid)
         	VALUES (?, ?, ?)
		`,
		f.Uuid, f.Person1Uuid, f.Person2Uuid,
	)

	return err
}

func (r *FamilyRepository) Fetch(uuid string) (*models.Family, error) {
	var f models.Family

	err := r.DB.QueryRow(
		`SELECT
    		uuid, person_1_uuid, person_2_uuid
         	FROM families
         	WHERE uuid = ?
		`,
		uuid,
	).Scan(&f.Uuid, &f.Person1Uuid, &f.Person2Uuid)

	if err != nil {
		return nil, err
	}
	return &f, nil
}

func (r *FamilyRepository) FetchAll() ([]models.Family, error) {
	var f []models.Family

	rows, err := r.DB.Query(`SELECT * FROM families`)
	if err != nil {
		return f, err
	}

	defer rows.Close()

	for rows.Next() {
		var row models.Family
		err := rows.Scan(&row.Uuid, &row.Person1Uuid, &row.Person2Uuid)

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

func (r *FamilyRepository) Update(f models.Family) error {
	_, err := r.DB.Exec(
		`
			UPDATE families
			SET person_1_uuid=?, person_2_uuid=?
			WHERE uuid=?
		`,
		f.Uuid, f.Person1Uuid, f.Person2Uuid,
	)

	return err
}

func (r *FamilyRepository) Delete(uuid string) error {
	_, err := r.DB.Exec(
		`DELETE FROM families WHERE uuid=?`,
		uuid,
	)

	return err
}
