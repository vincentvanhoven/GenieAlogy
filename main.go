package main

import (
	"embed"
	"log"
	"runtime"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/menu"
	"github.com/wailsapp/wails/v2/pkg/menu/keys"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	rt "github.com/wailsapp/wails/v2/pkg/runtime"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := NewApp()

	// Configure the menus
	AppMenu := menu.NewMenu()
	if runtime.GOOS == "darwin" {
		// On macOS platform, this must be done right after `NewMenu()`
		AppMenu.Append(menu.AppMenu())
	}

	FileMenu := AppMenu.AddSubmenu("File")
	FileMenu.AddText("Open", keys.CmdOrCtrl("o"), func(_ *menu.CallbackData) {
		saveFile, err := app.LoadFile()
		if err != nil {
			log.Default().Println(err)
		}

		if saveFile != nil {
			rt.EventsEmit(app.ctx, "onSaveFileLoaded", saveFile)
		}
	})
	FileMenu.AddSeparator()
	FileMenu.AddText("Save", keys.CmdOrCtrl("s"), func(_ *menu.CallbackData) {
		rt.EventsEmit(app.ctx, "onSaveRequested")
	})

	if runtime.GOOS == "darwin" {
		// On macOS platform, EditMenu should be appended to enable Cmd+C, Cmd+V, Cmd+Z... shortcuts
		AppMenu.Append(menu.EditMenu())
	}

	// Create the application
	err := wails.Run(&options.App{
		Title:  "GenieAlogy",
		Width:  1024,
		Height: 768,
		Menu:   AppMenu,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}

}
