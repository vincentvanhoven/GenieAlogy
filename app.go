package main

import (
	"GenieAlogy/repositories"
	"context"
	"database/sql"
	_ "database/sql"
	"log"

	"GenieAlogy/models"

	"github.com/pressly/goose/v3"
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

	if err := goose.Up(db, "database/migrations"); err != nil {
		log.Fatal(err)
	}
}

func InsertDataIntoDatabase(db *sql.DB, saveFile models.SaveFile) {
	personRepo := repositories.PersonRepository{db}

	for _, person := range saveFile.People {
		err := personRepo.Insert(person)

		if err != nil {
			log.Fatal(err)
		}
	}

	familyRepo := repositories.FamilyRepository{db}

	for _, family := range saveFile.Families {
		err := familyRepo.Insert(family)

		if err != nil {
			log.Fatal(err)
		}
	}
}

func (a *App) LoadFile() (*models.SaveFile, error) {
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

	db, _ := sql.Open("sqlite", path)
	defer db.Close()

	familyRepo := repositories.FamilyRepository{db}
	personRepo := repositories.PersonRepository{db}

	families, err := familyRepo.FetchAll()
	if err != nil {
		log.Fatal(err)
	}

	people, err := personRepo.FetchAll()
	if err != nil {
		log.Fatal(err)
	}

	saveFile := models.SaveFile{
		people,
		families,
	}

	// Keep track of open file/project path
	a.openFilePath = path

	return &saveFile, nil
}

func (a *App) SaveFile(saveFile models.SaveFile) error {
	// Open file picker
	path, err := runtime.SaveFileDialog(a.ctx, runtime.SaveDialogOptions{
		Filters: []runtime.FileFilter{
			{
				DisplayName: "Geniealogy files",
				Pattern:     "*.geniealogy",
			},
		},
	})
	if err != nil {
		return err
	}

	db, _ := sql.Open("sqlite", path)
	defer db.Close()

	MigrateDatabase(db)
	InsertDataIntoDatabase(db, saveFile)

	return nil
}
