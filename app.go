package main

import (
	"GenieAlogy/repositories"
	"context"
	"log"
	"strings"

	"GenieAlogy/models"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

const GENIEALOGY_VERSION = "0.0.2"

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
				Pattern:     "*.lamp",
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

func (a *App) NewFile() (*models.SaveFile, error) {
	// Open file picker
	path, err := runtime.SaveFileDialog(a.ctx, runtime.SaveDialogOptions{
		Filters: []runtime.FileFilter{
			{
				DisplayName: "Geniealogy files",
				Pattern:     "*.lamp",
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

	if !strings.HasSuffix(path, ".lamp") {
		path += ".lamp"
	}

	err = repositories.DatabaseRepo.Create(path)
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

func (a *App) UpdatePerson(person models.Person) error {
	err := repositories.PersonRepo.Update(person)
	return err
}

func (a *App) RemovePerson(person models.Person) (*models.Person, error) {
	families, err := repositories.FamilyRepo.FetchForPerson(person)

	if err != nil {
		return nil, err
	}

	hasDescendants := false

	// First, a soft-run to see if the Person has descendants.
	for _, family := range families {
		people, err := repositories.PersonRepo.FetchForFamily(family)
		if err != nil {
			return nil, err
		}

		if len(people) > 0 {
			hasDescendants = true
		}
	}

	if hasDescendants {
		// To prevent breaking the tree when there are descendants, anonymize the Person instead
		err = repositories.PersonRepo.Anonimize(person)
		if err != nil {
			return nil, err
		}

		p, err := repositories.PersonRepo.Fetch(*person.Id)
		if err != nil {
			return p, err
		}

		// Return the anonymized Person
		return p, nil
	} else {
		// Iterate familes
		for _, family := range families {
			err := repositories.FamilyRepo.Delete(*family.Id)
			if err != nil {
				return nil, err
			}
		}

		err = repositories.PersonRepo.Delete(*person.Id)
		if err != nil {
			return nil, err
		}

		return nil, nil
	}
}

func (a *App) AddFamily(family models.Family) (*models.Family, error) {
	id, err := repositories.FamilyRepo.Create(family)

	if err != nil {
		return nil, err
	}

	f, err := repositories.FamilyRepo.Fetch(*id)
	if err != nil {
		log.Fatal(err)
	}

	return f, nil
}

func (a *App) UpdateFamily(family models.Family) error {
	err := repositories.FamilyRepo.Update(family)
	return err
}

func (a *App) RemoveFamily(family models.Family) error {
	err := repositories.PersonRepo.ClearFamily(*family.Id)

	if err != nil {
		return err
	}

	err = repositories.FamilyRepo.Delete(*family.Id)

	if err != nil {
		return err
	}

	return nil
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
