package handlers

import (
	"kurohelper-api/middlware"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/peter910820/kurohelper-db/repository"
	"github.com/sirupsen/logrus"
)

// internal environment tokens generate handler
func TokensGenerateHandler(c *fiber.Ctx) error {
	id := uuid.New()

	// 寫到db(目前過期時間預設都是無限，因為只有內部使用)
	err := repository.CreateWebAPIToken(id.String(), 0)
	if err != nil {
		logrus.Error(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err,
		})
	}

	// 寫到快取
	middlware.VaildToken[id.String()] = struct{}{}

	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"message": "token generated successfully",
		"token":   id.String(),
	})
}
