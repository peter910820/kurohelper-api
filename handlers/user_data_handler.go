package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/peter910820/kurohelper-db/repository"
	"github.com/sirupsen/logrus"
)

func GetUserDataHandler(c *fiber.Ctx) error {
	// URL decoding
	id := c.Query("id")

	userGames, err := repository.GetUserData(id)
	if err != nil {
		logrus.Error(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "search successfully",
		"data":    userGames,
	})
}
