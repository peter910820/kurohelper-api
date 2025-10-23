package routes

import (
	"kurohelper-api/handlers"

	"github.com/gofiber/fiber/v2"
)

func TokenRouter(apiGroup fiber.Router) {
	tokenGroup := apiGroup.Group("/tokens")

	tokenGroup.Get("/generate", func(c *fiber.Ctx) error {
		return handlers.TokensGenerateHandler(c)
	})
}
