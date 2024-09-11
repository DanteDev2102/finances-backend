package middlewares

import (
	"github.com/gofiber/fiber/v2"
	monitor "github.com/gofiber/fiber/v2/middleware/monitor"
)

var Monitor fiber.Handler = monitor.New(monitor.Config{Title: "Metricts Page"})
