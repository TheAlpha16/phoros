package middleware

import (
	"strconv"
	"time"

	"github.com/TheAlpha16/phoros/src/config"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

func CheckTime() fiber.Handler {
	return func(c *fiber.Ctx) error {
		if config.EVENT_START == "" || config.EVENT_END == "" {
			return c.Next()
		}

		var startTime int64
		var endTime int64
		var err error

		startTime, err = strconv.ParseInt(config.EVENT_START, 10, 64) 
		if err == nil {
			return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{"status": "failure", "message": "invalid event start time"})
		}

		endTime, err = strconv.ParseInt(config.EVENT_END, 10, 64) 
		if err == nil {
			return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{"status": "failure", "message": "invalid event end time"})
		}

		if time.Now().Before(time.Unix(startTime, 0)) {
			return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{"status": "failure", "message": "event has not started yet"})
		}

		if time.Now().After(time.Unix(endTime, 0)) && config.POST_EVENT == "false" {
			return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{"status": "failure", "message": "event has ended"})
		}

		return c.Next()
	}
}

func CheckToken() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{
			Key:    []byte(config.SESSION_SECRET),
			JWTAlg: jwtware.HS256,
		},
		TokenLookup: "cookie:token", 
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"status":  "failure",
				"message": "invalid or expired session token",
			})
		},
	})
}
