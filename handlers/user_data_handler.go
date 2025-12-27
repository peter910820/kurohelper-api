package handlers

import (
	"github.com/gofiber/fiber/v2"
	kurohelperdb "github.com/peter910820/kurohelper-db/v2"
	"github.com/sirupsen/logrus"
)

func GetUserHasPlayedHandler(c *fiber.Ctx) error {
	// URL decoding
	id := c.Query("id")

	userHasPlayed, err := kurohelperdb.SelectUserHasPlayed(id)
	if err != nil {
		logrus.Error(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "search successfully",
		"data":    userHasPlayed,
	})
}
