package db

import (
	"fmt"
	"log"

	enviroment "finances/internal/config"
	entities "finances/internal/database/entities"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase() *gorm.DB {
	env := enviroment.Enviroments()

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		env.HostDB, env.UserDB, env.PassDB, env.NameDB, env.PortDB,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	return db
}

func SyncEntities() {
	db := ConnectDatabase()

	db.AutoMigrate(&entities.Rate{})
	db.AutoMigrate(&entities.Account{})
	db.AutoMigrate(&entities.Budget{})
	db.AutoMigrate(&entities.Operation{})
	db.AutoMigrate(&entities.Currency{})
	db.AutoMigrate(&entities.User{})
}
