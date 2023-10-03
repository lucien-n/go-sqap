package router

import (
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
)

func CreateRouter() *fiber.App {
	app := fiber.New(fiber.Config{
		JSONEncoder:   json.Marshal,
		JSONDecoder:   json.Unmarshal,
		AppName:       "SQAP",
		ServerHeader:  "Fiber",
		CaseSensitive: true,
	})

	return app
}
