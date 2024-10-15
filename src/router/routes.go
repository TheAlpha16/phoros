package router

import (
	"github.com/TheAlpha16/phoros/src/handler"
	"github.com/TheAlpha16/phoros/src/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/ping", handler.Ping)

	files := app.Group("/files", middleware.CheckTime(), middleware.CheckToken())
	files.Get("/:file", handler.GetFile)
}
