package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/sanjivyash/AuthAPI/config"
	"github.com/sanjivyash/AuthAPI/router/api"
)

func main() {
	// instantiate a server
	app := fiber.New()

	// setup middleware
	app.Use(logger.New())

	// loading the environment variables
	port := config.Config("PORT")
	endpoint := config.Config("ENDPOINT")

	// setup the route handler
	api.Router(app, endpoint)

	// listen to endpoints
	log.Fatal(app.Listen(port))
}
