package repositories_test

import (
	"GenieAlogy/database/seeders"
	"GenieAlogy/models"
	"GenieAlogy/repositories"
	"testing"

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
		nil,
		"male",
		seeders.Strptr("John"),
		seeders.Strptr("Doe"),
		seeders.Strptr("1950-01-01"),
		seeders.Strptr("New York"),
		nil,
		nil,
		0,
		0,
		seeders.Strptr("2020-01-01"),
		seeders.Strptr("New Jersey"),
	}

	createdId, err := repositories.PersonRepo.Create(personCreateData)

	if err != nil {
		t.Fatal(err)
	}

	// Save the newly created ID to the test data
	personCreateData.Id = createdId

	person, err := repositories.PersonRepo.Fetch(*personCreateData.Id)

	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, personCreateData, *person)
}

func TestPersonRepository_Update_Fetch(t *testing.T) {
	testSetup(t)

	personCreateData := models.Person{
		nil,
		"male",
		seeders.Strptr("John"),
		seeders.Strptr("Doe"),
		seeders.Strptr("1950-01-01"),
		seeders.Strptr("New York"),
		nil,
		nil,
		0,
		0,
		seeders.Strptr("2020-01-01"),
		seeders.Strptr("New Jersey"),
	}

	createdId, err := repositories.PersonRepo.Create(personCreateData)

	if err != nil {
		t.Fatal(err)
	}

	// Save the newly created ID to the test data
	personCreateData.Id = createdId

	personUpdateData := models.Person{
		personCreateData.Id,
		"male",
		seeders.Strptr("James"),
		seeders.Strptr("Johnson"),
		seeders.Strptr("1951-01-01"),
		seeders.Strptr("Los Angeles"),
		nil,
		nil,
		1,
		2,
		seeders.Strptr("2021-01-01"),
		seeders.Strptr("Las Vegas"),
	}

	err = repositories.PersonRepo.Update(personUpdateData)
	if err != nil {
		t.Fatal(err)
	}

	person, err := repositories.PersonRepo.Fetch(*personCreateData.Id)

	if err != nil {
		t.Fatal(err)
	}

	assert.NotEqual(t, personCreateData, *person)
	assert.Equal(t, personUpdateData, *person)
}

func TestPersonRepository_Delete_Fetch(t *testing.T) {
	testSetup(t)

	personCreateData := models.Person{
		nil,
		"male",
		seeders.Strptr("John"),
		seeders.Strptr("Doe"),
		seeders.Strptr("1950-01-01"),
		seeders.Strptr("New York"),
		nil,
		nil,
		0,
		0,
		seeders.Strptr("2020-01-01"),
		seeders.Strptr("New Jersey"),
	}

	createdId, err := repositories.PersonRepo.Create(personCreateData)

	if err != nil {
		t.Fatal(err)
	}

	// Save the newly created ID to the test data
	personCreateData.Id = createdId

	err = repositories.PersonRepo.Delete(*personCreateData.Id)
	if err != nil {
		t.Fatal(err)
	}

	person, err := repositories.PersonRepo.Fetch(*personCreateData.Id)

	if err == nil || person != nil {
		t.Fatal("person was not deleted")
	}
}

func TestPersonRepository_ClearFamily(t *testing.T) {
	testSetup(t)

	var createdPersonIds []int

	for _ = range 5 {
		personCreateData := models.Person{
			nil,
			"male",
			seeders.Strptr("John"),
			seeders.Strptr("Doe"),
			seeders.Strptr("1950-01-01"),
			seeders.Strptr("New York"),
			seeders.Intptr(123),
			nil,
			0,
			0,
			seeders.Strptr("2020-01-01"),
			seeders.Strptr("New Jersey"),
		}

		createdId, err := repositories.PersonRepo.Create(personCreateData)

		if err != nil {
			t.Fatal(err)
		}

		createdPersonIds = append(createdPersonIds, *createdId)
	}

	for _, id := range createdPersonIds {
		person, err := repositories.PersonRepo.Fetch(id)
		if err != nil {
			t.Fatal(err)
		}

		if *person.FamilyId != 123 {
			t.Fatal("Expected FamilyId 123, got ", person.FamilyId)
		}
	}

	err := repositories.PersonRepo.ClearFamily(123)
	if err != nil {
		return
	}

	for _, id := range createdPersonIds {
		person, err := repositories.PersonRepo.Fetch(id)
		if err != nil {
			t.Fatal(err)
		}

		if person.FamilyId != nil {
			t.Fatal("Expected FamilyId nil, got ", *person.FamilyId)
		}
	}
}
