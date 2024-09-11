package middlewares

import (
	"github.com/gofiber/fiber/v2"
	cors "github.com/gofiber/fiber/v2/middleware/cors"
)

var Cors fiber.Handler = cors.New(cors.Config{})
