package seeders

import (
	"GenieAlogy/models"
	"GenieAlogy/repositories"
	"database/sql"
	"log"

	"github.com/google/uuid"
)

func RunPeopleSeeder() {
	persons := []models.Person{
		{
			uuid.New().String(),
			"male",
			sql.NullString{"John", true},
			sql.NullString{"Doe", true},
			sql.NullString{"1950-01-01", true},
			sql.NullString{"New York", true},
			sql.NullString{"", false},
			sql.NullString{"", false},
			models.Position{X: 0, Y: 0},
		},
		{
			uuid.New().String(),
			"male",
			sql.NullString{"John", true},
			sql.NullString{"Doe", true},
			sql.NullString{"1950-01-01", true},
			sql.NullString{"New York", true},
			sql.NullString{"", false},
			sql.NullString{"", false},
			models.Position{X: 0, Y: 0},
		},
		{
			uuid.New().String(),
			"male",
			sql.NullString{"John", true},
			sql.NullString{"Doe", true},
			sql.NullString{"1950-01-01", true},
			sql.NullString{"New York", true},
			sql.NullString{"", false},
			sql.NullString{"", false},
			models.Position{X: 0, Y: 0},
		},
		{
			uuid.New().String(),
			"female",
			sql.NullString{"Jane", true},
			sql.NullString{"Doe", true},
			sql.NullString{"1950-01-01", true},
			sql.NullString{"New York", true},
			sql.NullString{"", false},
			sql.NullString{"", false},
			models.Position{X: 0, Y: 0},
		},
		{
			uuid.New().String(),
			"female",
			sql.NullString{"Jane", true},
			sql.NullString{"Doe", true},
			sql.NullString{"1950-01-01", true},
			sql.NullString{"New York", true},
			sql.NullString{"", false},
			sql.NullString{"", false},
			models.Position{X: 0, Y: 0},
		},
		{
			uuid.New().String(),
			"female",
			sql.NullString{"Jane", true},
			sql.NullString{"Doe", true},
			sql.NullString{"1950-01-01", true},
			sql.NullString{"New York", true},
			sql.NullString{"", false},
			sql.NullString{"", false},
			models.Position{X: 0, Y: 0},
		},
	}

	for _, p := range persons {
		err := repositories.PersonRepo.Create(p)

		if err != nil {
			log.Fatal(err)
		}
	}
}
