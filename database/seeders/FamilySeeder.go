package seeders

import (
	"GenieAlogy/models"
	"GenieAlogy/repositories"
	"database/sql"
	"fmt"
	"log"
	"strings"

	"github.com/google/uuid"
)

func RunFamilySeeder() {
	people, _ := repositories.PersonRepo.FetchAll()
	grouped := make(map[string][]models.Person)

	for _, p := range people {
		grouped[p.Sex] = append(grouped[p.Sex], p)
	}

	families := []models.Family{
		{
			uuid.New().String(),
			sql.NullString{grouped["male"][0].Uuid, true},
			sql.NullString{grouped["female"][0].Uuid, true},
		},
		{
			uuid.New().String(),
			sql.NullString{grouped["male"][1].Uuid, true},
			sql.NullString{grouped["female"][1].Uuid, true},
		},
	}

	var placeholders []string
	var args []interface{}

	for _, f := range families {
		placeholders = append(placeholders, "(?, ?, ?)")
		args = append(args, f.Uuid, f.Person1Uuid, f.Person2Uuid)
	}

	query := fmt.Sprintf(
		`INSERT INTO families (uuid, person_1_uuid, person_2_uuid) VALUES %s`,
		strings.Join(placeholders, ", "),
	)

	_, err := repositories.DatabaseRepo.DB.Exec(query, args...)
	if err != nil {
		log.Fatal(err)
	}
}
