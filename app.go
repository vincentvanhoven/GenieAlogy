package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"os"

	_ "database/sql"

	"github.com/pressly/goose/v3"
	_ "github.com/pressly/goose/v3"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	_ "modernc.org/sqlite"
)

const GENIEALOGY_VERSION = "0.0.1"

type App struct {
	ctx          context.Context
	openFilePath string
}

func NewApp() *App {
	return &App{}
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func MigrateDatabase(db *sql.DB) {
	if err := goose.SetDialect("sqlite3"); err != nil {
		log.Fatal(err)
	}

	if err := goose.Up(db, "migrations"); err != nil {
		log.Fatal(err)
	}
}

func InsertDataIntoDatabase(db *sql.DB, saveFile SaveFile) {
	for _, person := range saveFile.People {
		_, err := db.Exec(
			`INSERT OR REPLACE INTO people (
       			uuid,
				sex,
				firstname,
				lastname,
				birthdate,
				birthplace,
				position_x,
				position_y
       		) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`,
			person.Uuid,
			person.Sex,
			person.Firstname,
			person.Lastname,
			person.Birthdate,
			person.Birthplace,
			person.Position.X,
			person.Position.Y,
		)

		if err != nil {
			log.Fatal(err)
		}
	}

	for _, family := range saveFile.Families {
		_, err := db.Exec(
			`INSERT OR REPLACE INTO families (uuid, person_1_uuid, person_2_uuid) VALUES (?, ?, ?)`,
			family.Uuid,
			family.Person1Uuid,
			family.Person2Uuid,
		)

		if err != nil {
			log.Fatal(err)
		}
	}

	for _, person := range saveFile.People {
		_, err := db.Exec(
			`UPDATE people SET family_uuid = ? WHERE uuid = ?`,
			person.FamilyUuid,
			person.Uuid,
		)

		if err != nil {
			log.Fatal(err)
		}
	}
}

func (a *App) LoadFile() (*SaveFile, error) {
	// Open file picker
	path, err := runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions{
		Filters: []runtime.FileFilter{
			{
				DisplayName: "Geniealogy files",
				Pattern:     "*.geniealogy",
			},
		},
	})
	if err != nil {
		return nil, err
	}

	// Create handle for picked file
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	// Read & JSON-decode the picked file's contents
	var saveFile SaveFile
	decoder := json.NewDecoder(file)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&saveFile); err != nil {
		return nil, err
	}

	// Keep track of open file/project path
	a.openFilePath = path

	return &saveFile, nil
}

func (a *App) SaveFile(saveFile SaveFile) error {
	db, _ := sql.Open("sqlite", "/Users/vincent/Code/GenieAlogy/tests/data/testdata2.geniealogy")
	defer db.Close()

	MigrateDatabase(db)
	InsertDataIntoDatabase(db, saveFile)

	return nil
}
