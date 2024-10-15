package models

import (
	"github.com/gofiber/fiber/v2"
)

type S3FS struct {}

func (s *S3FS) Authenticate() error {
	return nil
}

func (s *S3FS) GetFile(c *fiber.Ctx, file string) error {
	return nil
}