package main

import (
	"fmt"
	"kurohelper-api/middlware"
	"kurohelper-api/routes"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"

	kurohelperdb "github.com/peter910820/kurohelper-db"
	"github.com/peter910820/kurohelper-db/models"
	"github.com/peter910820/kurohelper-db/repository"
)

func init() {
	// init logrus settings
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:   true,
		FullTimestamp: true,
	})
	logrus.SetLevel(logrus.DebugLevel)
	// init env file
	err := godotenv.Load()
	if err != nil {
		logrus.Fatalf(".env file load error: %v", err)
	}
}

func main() {
	config := models.Config{
		DBOwner:    os.Getenv("DB_OWNER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
		DBPort:     os.Getenv("DB_PORT"),
	}

	kurohelperdb.InitDsn(config)
	kurohelperdb.Migration() // 選填

	initTokenCache()

	app := fiber.New()

	// api route group
	apiGroup := app.Group("/api") // main api route group

	// site route group
	routes.TokenRouter(apiGroup)
	routes.UserDataRouter(apiGroup)

	logrus.Fatal(app.Listen(fmt.Sprintf("127.0.0.1:%s", os.Getenv("PRODUCTION_PORT"))))
}

func initTokenCache() {
	webAPIToken, err := repository.GetWebAPIToken()
	if err != nil {
		logrus.Fatal(err)
	}

	for _, t := range webAPIToken {
		middlware.VaildToken[t.ID] = struct{}{}
	}
}
