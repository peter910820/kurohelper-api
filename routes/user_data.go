package routes

import (
	"kurohelper-api/handlers"
	"kurohelper-api/middlware"

	"github.com/gofiber/fiber/v2"
)

func UserDataRouter(apiGroup fiber.Router) {
	userDataGroup := apiGroup.Group("/userdata")

	// 獲取指定使用者全部的遊玩資料
	userDataGroup.Get("/", middlware.TokenAuth(), func(c *fiber.Ctx) error {
		return handlers.GetUserDataHandler(c)
	})
}
