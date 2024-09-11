package middlewares

import (
	"github.com/gofiber/fiber/v2"
	helmet "github.com/gofiber/fiber/v2/middleware/helmet"
)

var Helmet fiber.Handler = helmet.New(helmet.Config{XSSProtection: "1"})
