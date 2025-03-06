package main

import (
	"log"

	"github.com/TheAlpha16/phoros/src/logs"
	"github.com/TheAlpha16/phoros/src/config"
	"github.com/TheAlpha16/phoros/src/router"
	"github.com/TheAlpha16/phoros/src/storage"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	// Initialize error logging
	logs.InitLogger()

	log.Println("phoros v1.1.1")

	// Authenticate to file storage
	if err := storage.InitStorage(); err != nil {
		panic("failed to authenticate to storage")
	}

	// Initialize *fiber.App
	app := fiber.New()
	app.Use(recover.New())            // Prevent process exit due to Fatal()
	router.SetupRoutes(app)           // Setup routing

	log.Fatal(app.Listen(config.APP_PORT))
}