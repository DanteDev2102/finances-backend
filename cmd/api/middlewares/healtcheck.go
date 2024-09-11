package middlewares

import (
	"github.com/gofiber/fiber/v2"
	healthcheck "github.com/gofiber/fiber/v2/middleware/healthcheck"
)

var Healthcheck fiber.Handler = healthcheck.New(healthcheck.Config{LivenessEndpoint: "/live", ReadinessEndpoint: "/ready"})
