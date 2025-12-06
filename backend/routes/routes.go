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

	// Protected routes
	auth := v1.Group("", middleware.Auth)

	// Auth
	auth.Get("/me", controllers.Me)

	// Lists
	auth.Post("/lists", controllers.CreateList)
	auth.Get("/lists", controllers.GetLists)
	auth.Get("/lists/:id", controllers.GetList)
	auth.Put("/lists/:id", controllers.UpdateList)
	auth.Delete("/lists/:id", controllers.DeleteList)
}
