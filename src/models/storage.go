package models

import (
	"github.com/gofiber/fiber/v2"
)

type Storage interface {
	Authenticate() error
	GetFile(c *fiber.Ctx, file string) error
	UploadFile(c *fiber.Ctx) error
}