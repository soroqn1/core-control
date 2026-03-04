package main

import (
	"core-control/internal/api"
	"core-control/internal/docker"
	"embed"
	"log"
	"net/http"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Initialize Docker service
	dockerService, err := docker.NewDockerService()
	if err != nil {
		log.Printf("Docker service initialization failed: %v", err)
	}
	defer func() {
		if dockerService != nil {
			dockerService.Close()
		}
	}()

	// Start HTTP API server in a separate goroutine
	if dockerService != nil {
		apiServer := api.NewServer(dockerService)
		mux := apiServer.SetupRoutes()
		go func() {
			log.Println("API Server started at http://localhost:8080")
			if err := http.ListenAndServe(":8080", mux); err != nil {
				log.Printf("API Server failed: %v", err)
			}
		}()
	}

	// Create an instance of the app structure
	app := NewApp()

	// Create application with options
	err = wails.Run(&options.App{
		Title:  "core-control",
		Width:  1024,
		Height: 768,
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
