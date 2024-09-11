package main

import (
	api "finances/cmd/api"
	enviroment "finances/internal/config"
	db "finances/internal/database"
)

func main() {
	app := api.App

	env := enviroment.Env

	db.SyncEntities()

	app.Listen(":" + env.Port)

}
