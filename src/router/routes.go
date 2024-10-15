package router

import (
	"github.com/TheAlpha16/phoros/src/handler"
	"github.com/TheAlpha16/phoros/src/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/ping", handler.Ping)

	admin := app.Group("/admin", middleware.AdminToken())
	admin.Post("/upload", handler.UploadFile)

	files := app.Group("/files", middleware.CheckTime(), middleware.CheckToken())
	files.Get("/:file", handler.GetFile)
}
