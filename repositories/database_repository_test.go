package repositories_test

import (
	"GenieAlogy/repositories"
	"os"
	"testing"
)

var testDatabaseFilePath = "./testing.geniealogy"

func TestDatabaseRepository_Create(t *testing.T) {
	ensureTestDatabaseDoesntExist(t)

	// Attempt DB creation
	err := repositories.DatabaseRepo.Create(testDatabaseFilePath)

	// Assert it did not fail
	if err != nil {
		t.Fatal(err)
	}

	// Assert the DB file exists now
	_, err = os.Stat(testDatabaseFilePath)
	if os.IsNotExist(err) {
		t.Fatal("DB was not created")
	}
}

func TestDatabaseRepository_Fetch_Update(t *testing.T) {
	ensureTestDatabaseDoesntExist(t)
	ensureTestDatabaseFileExists(t)

	err := repositories.DatabaseRepo.Fetch(testDatabaseFilePath)
	if err != nil {
		t.Fatal(err)
	}

	if repositories.DatabaseRepo.DB == nil {
		t.Fatal("DB was not opened")
	}

	err = repositories.DatabaseRepo.Update()
	if err != nil {
		t.Fatal(err)
	}
}

func ensureTestDatabaseFileExists(t *testing.T) {
	_, err := os.Stat(testDatabaseFilePath)
	if os.IsNotExist(err) {
		err := repositories.DatabaseRepo.Create(testDatabaseFilePath)
		if err != nil {
			t.Fatal(err)
		}

		err = repositories.DatabaseRepo.DB.Close()
		if err != nil {
			t.Fatal(err)
		}
	}
}

func ensureTestDatabaseDoesntExist(t *testing.T) {
	_, err := os.Stat(testDatabaseFilePath)
	if !os.IsNotExist(err) {
		err := os.Remove(testDatabaseFilePath)

		if err != nil {
			t.Fatal("Could not remove old test file")
		}
	}
}

func testSetup(t *testing.T) {
	// Remove old lingering DB file
	ensureTestDatabaseDoesntExist(t)
	// Create new DB file
	ensureTestDatabaseFileExists(t)

	err := repositories.DatabaseRepo.Fetch(testDatabaseFilePath)
	if err != nil {
		t.Fatal(err)
	}
}
