package handlers

import (
	"net/url"

	"github.com/gofiber/fiber/v2"
	"github.com/peter910820/kurohelper-db/repository"
	"github.com/sirupsen/logrus"
)

func GetUserDataHandler(c *fiber.Ctx) error {
	// URL decoding
	id, err := url.QueryUnescape(c.Params("id"))
	if err != nil {
		logrus.Error(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err,
		})
	}

	userGames, err := repository.GetUserData(id)
	if err != nil {
		logrus.Error(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err,
		})
	}

	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"message": "search successfully",
		"data":    userGames,
	})
}
