package repositories_test

import (
	"GenieAlogy/database/seeders"
	"GenieAlogy/models"
	"GenieAlogy/repositories"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFamilyRepository_FetchAll(t *testing.T) {
	testSetup(t)

	seeders.RunPeopleSeeder()
	seeders.RunFamilySeeder()

	families, err := repositories.FamilyRepo.FetchAll()
	if err != nil {
		t.Fatal(err)
	}

	if len(families) != 2 {
		t.Fatal("Expected 2 families, got ", len(families))
	}
}

func TestFamilyRepository_Create_Fetch(t *testing.T) {
	testSetup(t)

	familyCreateData := models.Family{
		nil,
		seeders.Intptr(1),
		seeders.Intptr(2),
		0,
		0,
	}

	createdId, err := repositories.FamilyRepo.Create(familyCreateData)

	if err != nil {
		t.Fatal(err)
	}

	// Save the newly created ID to the test data
	familyCreateData.Id = createdId

	family, err := repositories.FamilyRepo.Fetch(*familyCreateData.Id)

	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, familyCreateData, *family)
}

func TestFamilyRepository_Update_Fetch(t *testing.T) {
	testSetup(t)

	familyCreateData := models.Family{
		nil,
		seeders.Intptr(0),
		seeders.Intptr(1),
		0,
		0,
	}

	createdId, err := repositories.FamilyRepo.Create(familyCreateData)

	if err != nil {
		t.Fatal(err)
	}

	// Save the newly created ID to the test data
	familyCreateData.Id = createdId

	familyUpdateData := models.Family{
		familyCreateData.Id,
		seeders.Intptr(2),
		seeders.Intptr(3),
		0,
		0,
	}

	err = repositories.FamilyRepo.Update(familyUpdateData)
	if err != nil {
		t.Fatal(err)
	}

	family, err := repositories.FamilyRepo.Fetch(*familyCreateData.Id)

	if err != nil {
		t.Fatal(err)
	}

	assert.NotEqual(t, familyCreateData, *family)
	assert.Equal(t, familyUpdateData, *family)
}

func TestFamilyRepository_Delete_Fetch(t *testing.T) {
	testSetup(t)

	familyCreateData := models.Family{
		nil,
		seeders.Intptr(1),
		seeders.Intptr(2),
		0,
		0,
	}

	createdId, err := repositories.FamilyRepo.Create(familyCreateData)

	if err != nil {
		t.Fatal(err)
	}

	// Save the newly created ID to the test data
	familyCreateData.Id = createdId

	err = repositories.FamilyRepo.Delete(*familyCreateData.Id)
	if err != nil {
		t.Fatal(err)
	}

	family, err := repositories.FamilyRepo.Fetch(*familyCreateData.Id)

	if err == nil || family != nil {
		t.Fatal("family was not deleted")
	}
}
