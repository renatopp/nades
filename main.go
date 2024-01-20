package main

import (
	"embed"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func GuaranteePath(path string) error {
	// Stat the path to check if it exists
	_, err := os.Stat(path)

	// If the path doesn't exist, create it
	if os.IsNotExist(err) {
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			return fmt.Errorf("failed to create directory: %v", err)
		}
	} else if err != nil {
		// Return an error if there's an issue other than the path not existing
		return fmt.Errorf("failed to check directory existence: %v", err)
	}

	return nil
}

func main() {

	happ := fiber.New()
	happ.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})
	happ.Post("/upload/:map/:variant/:type", func(c fiber.Ctx) error {
		file, err := c.FormFile("file")
		if err != nil {
			return err
		}

		path := fmt.Sprintf("./.screenshots/%s/%s", c.Params("map"), c.Params("variant"))
		destination := fmt.Sprintf("./.screenshots/%s/%s/%s", c.Params("map"), c.Params("variant"), c.Params("type"))
		if err := GuaranteePath(path); err != nil {
			return err
		}

		if err := c.SaveFile(file, destination); err != nil {
			return err
		}

		c.SendStatus(200)
		return c.SendString("ok")
	})
	happ.Static("/static", "./.screenshots/")
	happ.Use(cors.New(cors.Config{
		AllowOriginsFunc: func(origin string) bool {
			return true
		},
	}))
	go happ.Listen(":26023")

	// Create an instance of the app structure
	app := NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "CAPIVARA NADELOG",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		OnShutdown:       app.shutdown,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
