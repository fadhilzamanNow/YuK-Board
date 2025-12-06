package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"yukboard/config"
	"yukboard/routes"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	config.ConnectDB()

	app := fiber.New()

	routes.Setup(app)

	log.Fatal(app.Listen(":3000"))
}
