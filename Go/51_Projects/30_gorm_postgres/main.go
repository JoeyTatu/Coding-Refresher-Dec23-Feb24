package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joeytatu/gorm-postgres/handlers"
	"github.com/joeytatu/gorm-postgres/models"
	"github.com/joeytatu/gorm-postgres/storage"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error reading .env file:", err)
	}

	config := &storage.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	}

	db, err := storage.NewConnection(config)
	if err != nil {
		log.Fatal("Error loading database:", err)
	}

	err = models.MigrateBooks(db)
	if err != nil {
		log.Fatal("Could not migrate database:", err)
	}

	app := fiber.New()
	r := handlers.Repository{
		DB: db,
	}
	r.SetUpRoutes(app)
	app.Listen(":8000")
}
