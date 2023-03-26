package main

import (
	"fmt"
	"go-fiber-crm-basic/database"
	"go-fiber-crm-basic/lead"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/lead", lead.GetLeads)
	app.Get("/api/v1/lead/:id", lead.GetLead)
	app.Post("/api/v1/lead", lead.NewLead)
	app.Delete("/api/v1/lead/:id", lead.DeleteLead)
}

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open(sqlite.Open("leads.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect DB")
	}
	fmt.Println("Connection opened to DB")
	database.DBConn.AutoMigrate(&lead.Lead{})
	fmt.Println("Database MIgrated")
}

func main() {
	app := fiber.New()
	initDatabase()
	setupRoutes(app)
	app.Use(cors.New(cors.Config{
		AllowOrigins: "localhost:3000",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))
	log.Fatal(app.Listen(":3000"))
}
