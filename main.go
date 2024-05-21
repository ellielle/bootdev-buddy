package main

import (
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/linux"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:       "Boot.dev Buddy",
		Width:       1024,
		Height:      768,
		StartHidden: true,
		// Set to true to remove border
		Frameless: false,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		// translucency is also controlled by the background in
		// the app's main style.css file
		BackgroundColour: &options.RGBA{R: 0, G: 0, B: 0, A: 100},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
		Linux: &linux.Options{
			WindowIsTranslucent: true,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
