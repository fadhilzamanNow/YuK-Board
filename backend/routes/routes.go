package routes

import (
	"github.com/gofiber/fiber/v2"
	"yukboard/controllers"
	"yukboard/middleware"
)

func Setup(app *fiber.App) {
	v1 := app.Group("/api/v1")

	// Auth (public)
	v1.Post("/register", controllers.Register)
	v1.Post("/login", controllers.Login)

	// Auth (protected)
	v1.Get("/me", middleware.Auth, controllers.Me)
}
