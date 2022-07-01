package main

import (
	"embed"

	"github.com/emuell/restic-browser/lib"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
)

//go:embed frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the restic app
	app := lib.NewResticBrowser()

	// Create wails application with options and bind the app to the frontend
	err := wails.Run(&options.App{
		Title:     "Restic Browser",
		Width:     1024,
		Height:    768,
		Assets:    assets,
		OnStartup: app.Startup,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err)
	}
}
