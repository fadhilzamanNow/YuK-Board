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

	// Invitations
	auth.Post("/lists/:id/invite", controllers.InviteUser)
	auth.Get("/invitations", controllers.GetMyInvitations)
	auth.Post("/invitations/:id/accept", controllers.AcceptInvitation)
	auth.Post("/invitations/:id/decline", controllers.DeclineInvitation)
	auth.Delete("/invitations/:id", controllers.CancelInvitation)

	// Tasks
	auth.Post("/lists/:listId/tasks", controllers.CreateTask)
	auth.Get("/lists/:listId/tasks", controllers.GetTasks)
	auth.Get("/lists/:listId/tasks/:taskId", controllers.GetTask)
	auth.Put("/lists/:listId/tasks/:taskId", controllers.UpdateTask)
	auth.Delete("/lists/:listId/tasks/:taskId", controllers.DeleteTask)
	auth.Patch("/lists/:listId/tasks/:taskId/status", controllers.UpdateTaskStatus)
}
