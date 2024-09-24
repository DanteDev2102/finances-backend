package api

import (
	middlewares "finances/cmd/api/middlewares"

	"github.com/gofiber/fiber/v2"
)

var App *fiber.App = fiber.New(fiber.Config{CaseSensitive: true, AppName: "Finances API v2.0.0"})
var RouterHandler fiber.Router = App.Group("/api/v2")

func init() {
	App.Use(middlewares.Healthcheck)
	App.Use(middlewares.Cors)
	App.Use(middlewares.Helmet)
	
	RouterHandler.Get("/monitor", middlewares.Monitor)
}
