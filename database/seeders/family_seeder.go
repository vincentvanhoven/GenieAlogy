package seeders

import (
	"GenieAlogy/models"
	"GenieAlogy/repositories"
	"log"

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
			Strptr(grouped["male"][0].Uuid),
			Strptr(grouped["female"][0].Uuid),
			0,
			0,
		},
		{
			uuid.New().String(),
			Strptr(grouped["male"][1].Uuid),
			Strptr(grouped["female"][1].Uuid),
			0,
			0,
		},
	}

	for _, f := range families {
		err := repositories.FamilyRepo.Create(f)
		if err != nil {
			log.Fatal(err)
		}
	}
}
