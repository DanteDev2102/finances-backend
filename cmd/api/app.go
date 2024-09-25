package api

import (
	middlewares "finances/cmd/api/middlewares"
	services "finances/cmd/api/services"
	models "finances/internal/models"

	"github.com/gofiber/fiber/v2"
)

var App *fiber.App = fiber.New(fiber.Config{CaseSensitive: true, AppName: "Finances API v2.0.0"})
var RouterHandler fiber.Router = App.Group("/api/v2")

func init() {
	App.Use(middlewares.Healthcheck)
	App.Use(middlewares.Cors)
	App.Use(middlewares.Helmet)
	
	RouterHandler.Get("/monitor", middlewares.Monitor)
	
	RouterHandler.Post("/auth/signup", func(c *fiber.Ctx) error {
		user := new(models.MUser)

		if err := c.BodyParser(user); err != nil {
			return err
		}

		newUser, err := services.CreateUserService(user)

		if err != nil {
			return err
		}

		return c.JSON(newUser)
	})
}
