package handler

import (
	"github.com/TheAlpha16/phoros/src/utils"
	"github.com/TheAlpha16/phoros/src/storage"

	"github.com/gofiber/fiber/v2"
)

func GetFile(c *fiber.Ctx) error {
	file := c.Params("file")
	if _, err := utils.ValidateFileName(file); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("file does not exist")
	}

	err := storage.Store.GetFile(c, file)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("file does not exist")
	}
	
	return nil
}