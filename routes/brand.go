package routes

import (
	"api/handlers"

	"github.com/gofiber/fiber/v2"
)

func BrandRouter(apiGroup fiber.Router) {
	userDataGroup := apiGroup.Group("/brand")

	userDataGroup.Get("/", func(c *fiber.Ctx) error {
		return handlers.SearchBrandHandler(c)
	})
}
