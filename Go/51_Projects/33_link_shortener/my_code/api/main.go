package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joeytatu/link-shortener/routes"
	"github.com/joho/godotenv"
)

func setUpRoutes(app *fiber.App) {
	app.Get("/:url", routes.ResolveURL)
	app.Post("/api/v1", routes.ShortenURL)
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("could not load .env file: ", err)
	}

	app := fiber.New()
	app.Use(logger.New())
	setUpRoutes(app)

	log.Fatal(app.Listen(os.Getenv("APP_PORT")))
}
