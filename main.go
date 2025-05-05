package main

import (
	"log"

	"github.com/cuffymate1/pos-api/config"
	"github.com/cuffymate1/pos-api/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("config/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config.ConnDb()

	app := fiber.New()

	routes.GetRoutes(app)

	if err := app.Listen(":3000"); err != nil {
		log.Fatal(err)
	}
}
