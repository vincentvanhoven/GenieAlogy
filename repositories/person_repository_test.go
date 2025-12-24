package repositories_test

import (
	"GenieAlogy/database/seeders"
	"GenieAlogy/models"
	"GenieAlogy/repositories"
	"database/sql"
	"testing"

	"github.com/google/uuid"
)

func TestPersonRepository_FetchAll(t *testing.T) {
	testSetup(t)

	seeders.RunPeopleSeeder()

	people, err := repositories.PersonRepo.FetchAll()
	if err != nil {
		t.Fatal(err)
	}

	if len(people) != 6 {
		t.Fatal("Expected 6 people, got ", len(people))
	}
}

func TestPersonRepository_Create_Fetch(t *testing.T) {
	testSetup(t)

	personCreateData := models.Person{
		uuid.New().String(),
		"male",
		sql.NullString{"John", true},
		sql.NullString{"Doe", true},
		sql.NullString{"1950-01-01", true},
		sql.NullString{"New York", true},
		sql.NullString{"", false},
		sql.NullString{"", false},
		models.Position{X: 0, Y: 0},
	}

	err := repositories.PersonRepo.Create(personCreateData)

	if err != nil {
		t.Fatal(err)
	}

	person, err := repositories.PersonRepo.Fetch(personCreateData.Uuid)

	if err != nil {
		t.Fatal(err)
	}

	if *person != personCreateData {
		t.Fatal("person could not be inserted")
	}
}

func TestPersonRepository_Update_Fetch(t *testing.T) {
	testSetup(t)

	personCreateData := models.Person{
		uuid.New().String(),
		"male",
		sql.NullString{"John", true},
		sql.NullString{"Doe", true},
		sql.NullString{"1950-01-01", true},
		sql.NullString{"New York", true},
		sql.NullString{"", false},
		sql.NullString{"", false},
		models.Position{X: 0, Y: 0},
	}

	err := repositories.PersonRepo.Create(personCreateData)

	if err != nil {
		t.Fatal(err)
	}

	personUpdateData := models.Person{
		personCreateData.Uuid,
		"male",
		sql.NullString{"Jane", true},
		sql.NullString{"Johnson", true},
		sql.NullString{"1951-01-01", true},
		sql.NullString{"Los Angeles", true},
		sql.NullString{"", false},
		sql.NullString{"", false},
		models.Position{X: 1, Y: 2},
	}

	err = repositories.PersonRepo.Update(personUpdateData)
	if err != nil {
		t.Fatal(err)
	}

	person, err := repositories.PersonRepo.Fetch(personCreateData.Uuid)

	if err != nil {
		t.Fatal(err)
	}

	if *person == personCreateData {
		t.Fatal("person was not updated")
	}

	if *person != personUpdateData {
		t.Fatal("person was updated, but not correctly")
	}
}

func TestPersonRepository_Delete_Fetch(t *testing.T) {
	testSetup(t)

	personCreateData := models.Person{
		uuid.New().String(),
		"male",
		sql.NullString{"John", true},
		sql.NullString{"Doe", true},
		sql.NullString{"1950-01-01", true},
		sql.NullString{"New York", true},
		sql.NullString{"", false},
		sql.NullString{"", false},
		models.Position{X: 0, Y: 0},
	}

	err := repositories.PersonRepo.Create(personCreateData)

	if err != nil {
		t.Fatal(err)
	}

	err = repositories.PersonRepo.Delete(personCreateData.Uuid)
	if err != nil {
		t.Fatal(err)
	}

	person, err := repositories.PersonRepo.Fetch(personCreateData.Uuid)

	if err == nil || person != nil {
		t.Fatal("person was not deleted")
	}
}
