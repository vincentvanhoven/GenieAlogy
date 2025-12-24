package main

import (
	"GenieAlogy/database/seeders"
	"GenieAlogy/repositories"
	"log"
)

func main() {
	err := repositories.DatabaseRepo.Fetch("test-database.geniealogy")
	if err != nil {
		log.Fatal(err)
	}

	err = repositories.DatabaseRepo.Update()
	if err != nil {
		log.Fatal(err)
	}

	seeders.RunPeopleSeeder()
	seeders.RunFamilySeeder()
}
