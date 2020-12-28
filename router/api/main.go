package api

import (
	"github.com/gofiber/fiber/v2"

	"github.com/sanjivyash/AuthAPI/middleware"
)

func Router(app *fiber.App, endpoint string) {
	api := app.Group("/" + endpoint)
	info := api.Group("/info", middleware.Auth)

	// we have a few routes
	api.Post("/signup", createUser)
	api.Post("/login", loginUser)
	api.Post("/delete", deleteUser)
	
	info.Get("/", getInfo)
}
