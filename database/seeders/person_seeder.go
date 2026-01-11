package seeders

import (
	"GenieAlogy/models"
	"GenieAlogy/repositories"
	"log"
)

func RunPeopleSeeder() {
	persons := []models.Person{
		{
			nil,
			"male",
			Strptr("John"),
			Strptr("Doe"),
			Strptr("1950-01-01"),
			Strptr("New York"),
			nil,
			nil,
			0,
			0,
		},
		{
			nil,
			"male",
			Strptr("John"),
			Strptr("Doe"),
			Strptr("1950-01-01"),
			Strptr("New York"),
			nil,
			nil,
			0,
			0,
		},
		{
			nil,
			"male",
			Strptr("John"),
			Strptr("Doe"),
			Strptr("1950-01-01"),
			Strptr("New York"),
			nil,
			nil,
			0,
			0,
		},
		{
			nil,
			"female",
			Strptr("Jane"),
			Strptr("Doe"),
			Strptr("1950-01-01"),
			Strptr("New York"),
			nil,
			nil,
			0,
			0,
		},
		{
			nil,
			"female",
			Strptr("Jane"),
			Strptr("Doe"),
			Strptr("1950-01-01"),
			Strptr("New York"),
			nil,
			nil,
			0,
			0,
		},
		{
			nil,
			"female",
			Strptr("Jane"),
			Strptr("Doe"),
			Strptr("1950-01-01"),
			Strptr("New York"),
			nil,
			nil,
			0,
			0,
		},
	}

	for _, p := range persons {
		_, err := repositories.PersonRepo.Create(p)

		if err != nil {
			log.Fatal(err)
		}
	}
}
