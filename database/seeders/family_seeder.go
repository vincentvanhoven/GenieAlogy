package seeders

import (
	"GenieAlogy/models"
	"GenieAlogy/repositories"
	"log"
)

func RunFamilySeeder() {
	people, _ := repositories.PersonRepo.FetchAll()
	grouped := make(map[string][]models.Person)

	for _, p := range people {
		grouped[p.Sex] = append(grouped[p.Sex], p)
	}

	families := []models.Family{
		{
			nil,
			grouped["male"][0].Id,
			grouped["female"][0].Id,
			0,
			0,
		},
		{
			nil,
			grouped["male"][1].Id,
			grouped["female"][1].Id,
			0,
			0,
		},
	}

	for _, f := range families {
		_, err := repositories.FamilyRepo.Create(f)
		if err != nil {
			log.Fatal(err)
		}
	}
}
