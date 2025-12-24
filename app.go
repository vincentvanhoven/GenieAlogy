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
	// Open file picker
	//path, err := runtime.SaveFileDialog(a.ctx, runtime.SaveDialogOptions{
	//	Filters: []runtime.FileFilter{
	//		{
	//			DisplayName: "Geniealogy files",
	//			Pattern:     "*.geniealogy",
	//		},
	//	},
	//})
	//if err != nil {
	//	return err
	//}

	for _, person := range saveFile.People {
		err := repositories.PersonRepo.Create(person)

		if err != nil {
			log.Fatal(err)
		}
	}

	for _, family := range saveFile.Families {
		err := repositories.FamilyRepo.Create(family)

		if err != nil {
			log.Fatal(err)
		}
	}

	return nil
}
