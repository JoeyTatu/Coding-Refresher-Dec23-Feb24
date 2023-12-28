package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	"github.com/joeytatu/go-fiber-crm-basic/database"
	"github.com/joeytatu/go-fiber-crm-basic/lead"
)

func setUpRoutes(app *fiber.App) {
	app.Get("/api/leads", lead.GetAllLeads)
	app.Get("/api/leads/lead/:id", lead.GetLeadById)
	app.Post("/api/leads/lead", lead.NewLead)
	app.Delete("/api/leads/lead/:id", lead.DeleteLead)
}

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "leads.db")
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}
	fmt.Println("Connected to database.")

	database.DBConn.AutoMigrate(&lead.Lead{})
	fmt.Println("Database migrated.")
}

func main() {
	app := fiber.New()
	initDatabase()
	setUpRoutes(app)
	app.Listen(3000)
	defer database.DBConn.Close()
}
