package main

import (
	enviroment "finances/internal/config"
	db "finances/internal/database"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	env := enviroment.Enviroments()

	db.SyncEntities()

	app.Listen(":" + env.Port)
}
