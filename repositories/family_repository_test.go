package repositories_test

import (
	"GenieAlogy/database/seeders"
	"GenieAlogy/models"
	"GenieAlogy/repositories"
	"database/sql"
	"testing"

	"github.com/google/uuid"
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
		sql.NullString{uuid.New().String(), true},
		sql.NullString{uuid.New().String(), true},
	}

	err := repositories.FamilyRepo.Create(familyCreateData)

	if err != nil {
		t.Fatal(err)
	}

	family, err := repositories.FamilyRepo.Fetch(familyCreateData.Uuid)

	if err != nil {
		t.Fatal(err)
	}

	if *family != familyCreateData {
		t.Fatal("family could not be inserted")
	}
}

func TestFamilyRepository_Update_Fetch(t *testing.T) {
	testSetup(t)

	familyCreateData := models.Family{
		uuid.New().String(),
		sql.NullString{uuid.New().String(), true},
		sql.NullString{uuid.New().String(), true},
	}

	err := repositories.FamilyRepo.Create(familyCreateData)

	if err != nil {
		t.Fatal(err)
	}

	familyUpdateData := models.Family{
		familyCreateData.Uuid,
		sql.NullString{uuid.New().String(), true},
		sql.NullString{uuid.New().String(), true},
	}

	err = repositories.FamilyRepo.Update(familyUpdateData)
	if err != nil {
		t.Fatal(err)
	}

	family, err := repositories.FamilyRepo.Fetch(familyCreateData.Uuid)

	if err != nil {
		t.Fatal(err)
	}

	if *family == familyCreateData {
		t.Fatal("family was not updated")
	}

	if *family != familyUpdateData {
		t.Fatal("family was updated, but not correctly")
	}
}

func TestFamilyRepository_Delete_Fetch(t *testing.T) {
	testSetup(t)

	familyCreateData := models.Family{
		uuid.New().String(),
		sql.NullString{uuid.New().String(), true},
		sql.NullString{uuid.New().String(), true},
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
