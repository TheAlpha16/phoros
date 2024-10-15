package handler

import (
	"github.com/TheAlpha16/phoros/src/storage"
	
	"github.com/gofiber/fiber/v2"
)

func UploadFile(c *fiber.Ctx) error {
	err := storage.Store.UploadFile(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("upload error")
	}
	return nil
}