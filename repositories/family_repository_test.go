package repositories_test

import (
	"GenieAlogy/database/seeders"
	"GenieAlogy/models"
	"GenieAlogy/repositories"
	"testing"

	"github.com/google/uuid"
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
		uuid.New().String(),
		seeders.Strptr(uuid.New().String()),
		seeders.Strptr(uuid.New().String()),
		0,
		0,
	}

	err := repositories.FamilyRepo.Create(familyCreateData)

	if err != nil {
		t.Fatal(err)
	}

	family, err := repositories.FamilyRepo.Fetch(familyCreateData.Uuid)

	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, familyCreateData, *family)
}

func TestFamilyRepository_Update_Fetch(t *testing.T) {
	testSetup(t)

	familyCreateData := models.Family{
		uuid.New().String(),
		seeders.Strptr(uuid.New().String()),
		seeders.Strptr(uuid.New().String()),
		0,
		0,
	}

	err := repositories.FamilyRepo.Create(familyCreateData)

	if err != nil {
		t.Fatal(err)
	}

	familyUpdateData := models.Family{
		familyCreateData.Uuid,
		seeders.Strptr(uuid.New().String()),
		seeders.Strptr(uuid.New().String()),
		0,
		0,
	}

	err = repositories.FamilyRepo.Update(familyUpdateData)
	if err != nil {
		t.Fatal(err)
	}

	family, err := repositories.FamilyRepo.Fetch(familyCreateData.Uuid)

	if err != nil {
		t.Fatal(err)
	}

	assert.NotEqual(t, familyCreateData, *family)
	assert.Equal(t, familyUpdateData, *family)
}

func TestFamilyRepository_Delete_Fetch(t *testing.T) {
	testSetup(t)

	familyCreateData := models.Family{
		uuid.New().String(),
		seeders.Strptr(uuid.New().String()),
		seeders.Strptr(uuid.New().String()),
		0,
		0,
	}

	err := repositories.FamilyRepo.Create(familyCreateData)

	if err != nil {
		t.Fatal(err)
	}

	err = repositories.FamilyRepo.Delete(familyCreateData.Uuid)
	if err != nil {
		t.Fatal(err)
	}

	family, err := repositories.FamilyRepo.Fetch(familyCreateData.Uuid)

	if err == nil || family != nil {
		t.Fatal("family was not deleted")
	}
}
