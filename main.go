package main

import (
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "nameCaller",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 0, G: 0, B: 0, A: 1},
		CSSDragProperty:   "--wails-draggable",
        CSSDragValue:      "drag",
		OnStartup:        app.startup,
		Frameless:       true,
		Bind: []interface{}{
			app,
		},
		Windows: &windows.Options{
			WebviewGpuIsDisabled: false,
			Theme:                windows.SystemDefault,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}