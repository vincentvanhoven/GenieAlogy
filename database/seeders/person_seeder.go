package seeders

import (
	"GenieAlogy/models"
	"GenieAlogy/repositories"
	"log"

	"github.com/google/uuid"
)

func RunPeopleSeeder() {
	persons := []models.Person{
		{
			uuid.New().String(),
			"male",
			Strptr("John"),
			Strptr("Doe"),
			Strptr("1950-01-01"),
			Strptr("New York"),
			nil,
			nil,
			models.Position{X: 0, Y: 0},
		},
		{
			uuid.New().String(),
			"male",
			Strptr("John"),
			Strptr("Doe"),
			Strptr("1950-01-01"),
			Strptr("New York"),
			nil,
			nil,
			models.Position{X: 0, Y: 0},
		},
		{
			uuid.New().String(),
			"male",
			Strptr("John"),
			Strptr("Doe"),
			Strptr("1950-01-01"),
			Strptr("New York"),
			nil,
			nil,
			models.Position{X: 0, Y: 0},
		},
		{
			uuid.New().String(),
			"female",
			Strptr("Jane"),
			Strptr("Doe"),
			Strptr("1950-01-01"),
			Strptr("New York"),
			nil,
			nil,
			models.Position{X: 0, Y: 0},
		},
		{
			uuid.New().String(),
			"female",
			Strptr("Jane"),
			Strptr("Doe"),
			Strptr("1950-01-01"),
			Strptr("New York"),
			nil,
			nil,
			models.Position{X: 0, Y: 0},
		},
		{
			uuid.New().String(),
			"female",
			Strptr("Jane"),
			Strptr("Doe"),
			Strptr("1950-01-01"),
			Strptr("New York"),
			nil,
			nil,
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
