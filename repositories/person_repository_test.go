package repositories_test

import (
	"GenieAlogy/database/seeders"
	"GenieAlogy/models"
	"GenieAlogy/repositories"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
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
		seeders.Strptr("John"),
		seeders.Strptr("Doe"),
		seeders.Strptr("1950-01-01"),
		seeders.Strptr("New York"),
		nil,
		nil,
		0,
		0,
	}

	err := repositories.PersonRepo.Create(personCreateData)

	if err != nil {
		t.Fatal(err)
	}

	person, err := repositories.PersonRepo.Fetch(personCreateData.Uuid)

	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, personCreateData, *person)
}

func TestPersonRepository_Update_Fetch(t *testing.T) {
	testSetup(t)

	personCreateData := models.Person{
		uuid.New().String(),
		"male",
		seeders.Strptr("John"),
		seeders.Strptr("Doe"),
		seeders.Strptr("1950-01-01"),
		seeders.Strptr("New York"),
		nil,
		nil,
		0,
		0,
	}

	err := repositories.PersonRepo.Create(personCreateData)

	if err != nil {
		t.Fatal(err)
	}

	personUpdateData := models.Person{
		personCreateData.Uuid,
		"male",
		seeders.Strptr("James"),
		seeders.Strptr("Johnson"),
		seeders.Strptr("1951-01-01"),
		seeders.Strptr("Los Angeles"),
		nil,
		nil,
		1,
		2,
	}

	err = repositories.PersonRepo.Update(personUpdateData)
	if err != nil {
		t.Fatal(err)
	}

	person, err := repositories.PersonRepo.Fetch(personCreateData.Uuid)

	if err != nil {
		t.Fatal(err)
	}

	assert.NotEqual(t, personCreateData, *person)
	assert.Equal(t, personUpdateData, *person)
}

func TestPersonRepository_Delete_Fetch(t *testing.T) {
	testSetup(t)

	personCreateData := models.Person{
		uuid.New().String(),
		"male",
		seeders.Strptr("John"),
		seeders.Strptr("Doe"),
		seeders.Strptr("1950-01-01"),
		seeders.Strptr("New York"),
		nil,
		nil,
		0,
		0,
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
