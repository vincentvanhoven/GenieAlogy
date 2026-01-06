package main

import (
	"GenieAlogy/repositories"
	"context"
	"errors"
	"log"

	"GenieAlogy/models"

	"github.com/wailsapp/wails/v2/pkg/runtime"
	"modernc.org/sqlite"
	sqlite3 "modernc.org/sqlite/lib"
	//"modernc.org/sqlite/lib"
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

	err = repositories.DatabaseRepo.Fetch(path)
	if err != nil {
		log.Fatal(err)
	}

	families, err := repositories.FamilyRepo.FetchAll()
	if err != nil {
		log.Fatal(err)
	}

	people, err := repositories.PersonRepo.FetchAll()
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
	var sqliteErr *sqlite.Error

	for _, person := range saveFile.People {
		err := repositories.PersonRepo.Create(person)

		if err != nil && errors.As(err, &sqliteErr) && sqliteErr.Code() == sqlite3.SQLITE_CONSTRAINT_PRIMARYKEY {
			err := repositories.PersonRepo.Update(person)

			if err != nil {
				log.Fatal("SaveFile -> person.update", err)
			}
		} else if err != nil {
			log.Fatal("SaveFile -> person.create", err)
		}
	}

	for _, family := range saveFile.Families {
		err := repositories.FamilyRepo.Create(family)

		if err != nil && errors.As(err, &sqliteErr) && sqliteErr.Code() == sqlite3.SQLITE_CONSTRAINT_PRIMARYKEY {
			err := repositories.FamilyRepo.Update(family)

			if err != nil {
				log.Fatal("SaveFile -> family.update", err)
			}
		} else if err != nil {
			log.Fatal("SaveFile -> family.create", err)
		}
	}

	return nil
}
