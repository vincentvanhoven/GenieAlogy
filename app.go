package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

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
	// Create handle for picked file
	file, err := os.Create(a.openFilePath)
	if err != nil {
		return err
	}

	defer file.Close()

	// JSON-encode & write the savefile to the open file/project path
	encoder := json.NewEncoder(file)
	if err := encoder.Encode(saveFile); err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
