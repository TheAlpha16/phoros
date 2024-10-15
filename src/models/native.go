package models

import (
	"fmt"

	"github.com/TheAlpha16/phoros/src/utils"
	"github.com/TheAlpha16/phoros/src/config"

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