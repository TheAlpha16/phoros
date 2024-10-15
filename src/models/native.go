package models

import (
	"fmt"
	"log"

	"github.com/TheAlpha16/phoros/src/config"
	"github.com/TheAlpha16/phoros/src/utils"

	"github.com/gofiber/fiber/v2"
)

type NativeFS struct {}

func (n *NativeFS) Authenticate() error {
	return nil
}

func (n *NativeFS) GetFile(c *fiber.Ctx, file string) error {
	filePath := fmt.Sprintf("%s/%s", config.NFS_PATH, file)

	if !utils.FileExists(filePath) {
		return c.Status(fiber.StatusNotFound).SendString("file not found")
	}

	c.Set(fiber.HeaderContentDisposition, fmt.Sprintf(`attachment; filename="%s"`, file))

	return c.SendFile(filePath)
}

func (n *NativeFS) UploadFile(c *fiber.Ctx) error {
	file, err := c.FormFile("file")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("file missing in the request")
	}

	sanitizedName, err := utils.ValidateFileName(file.Filename)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("no file name provided")
	}

	filePath := fmt.Sprintf("%s/%s", config.NFS_PATH, sanitizedName)

	if utils.FileExists(filePath) {
		return c.Status(fiber.StatusConflict).SendString(fmt.Sprintf("/files/%s", sanitizedName))
	}

	if err := c.SaveFile(file, filePath); err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).SendString("failed to save file")
	}

	return c.Status(fiber.StatusCreated).SendString(fmt.Sprintf("/files/%s", sanitizedName))
}