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
			ptr("John"),
			ptr("Doe"),
			ptr("1950-01-01"),
			ptr("New York"),
			ptr(""),
			ptr(""),
			models.Position{X: 0, Y: 0},
		},
		{
			uuid.New().String(),
			"male",
			ptr("John"),
			ptr("Doe"),
			ptr("1950-01-01"),
			ptr("New York"),
			ptr(""),
			ptr(""),
			models.Position{X: 0, Y: 0},
		},
		{
			uuid.New().String(),
			"male",
			ptr("John"),
			ptr("Doe"),
			ptr("1950-01-01"),
			ptr("New York"),
			ptr(""),
			ptr(""),
			models.Position{X: 0, Y: 0},
		},
		{
			uuid.New().String(),
			"female",
			ptr("Jane"),
			ptr("Doe"),
			ptr("1950-01-01"),
			ptr("New York"),
			ptr(""),
			ptr(""),
			models.Position{X: 0, Y: 0},
		},
		{
			uuid.New().String(),
			"female",
			ptr("Jane"),
			ptr("Doe"),
			ptr("1950-01-01"),
			ptr("New York"),
			ptr(""),
			ptr(""),
			models.Position{X: 0, Y: 0},
		},
		{
			uuid.New().String(),
			"female",
			ptr("Jane"),
			ptr("Doe"),
			ptr("1950-01-01"),
			ptr("New York"),
			ptr(""),
			ptr(""),
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
