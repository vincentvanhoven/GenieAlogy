package main

import (
	"GenieAlogy/repositories"
	"context"
	"log"

	"GenieAlogy/models"

	"github.com/wailsapp/wails/v2/pkg/runtime"
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

	// Empty response if no file was chosen
	if path == "" {
		return nil, nil
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

func (a *App) AddPerson(position_x int, position_y int) (*models.Person, error) {
	createdId, err := repositories.PersonRepo.Create(models.Person{
		Sex:       "male",
		PositionX: position_x,
		PositionY: position_y,
	})

	if err != nil {
		return nil, err
	}

	person, err := repositories.PersonRepo.Fetch(*createdId)
	return person, err
}

func (a *App) SaveFile(saveFile models.SaveFile) error {
	for _, person := range saveFile.People {
		if person.Id == nil {
			_, err := repositories.PersonRepo.Create(person)

			if err != nil {
				log.Fatal("SaveFile -> person.create", err)
			}
		} else {
			err := repositories.PersonRepo.Update(person)

			if err != nil {
				log.Fatal("SaveFile -> person.update", err)
			}
		}
	}

	for _, family := range saveFile.Families {
		if family.Id == nil {
			_, err := repositories.FamilyRepo.Create(family)

			if err != nil {
				log.Fatal("SaveFile -> family.create", err)
			}
		} else {
			err := repositories.FamilyRepo.Update(family)

			if err != nil {
				log.Fatal("SaveFile -> family.update", err)
			}
		}
	}

	return nil
}
