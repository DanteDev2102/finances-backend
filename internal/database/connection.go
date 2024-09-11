package db

import (
	"fmt"
	"log"

	enviroment "finances/internal/config"
	entities "finances/internal/database/entities"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	env *enviroment.TEnv = enviroment.Env
	DB  *gorm.DB
)

func init() {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		env.HostDB, env.UserDB, env.PassDB, env.NameDB, env.PortDB,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	DB = db
}

func SyncEntities() {
	DB.AutoMigrate(&entities.Rate{})
	DB.AutoMigrate(&entities.Account{})
	DB.AutoMigrate(&entities.Budget{})
	DB.AutoMigrate(&entities.Operation{})
	DB.AutoMigrate(&entities.Currency{})
	DB.AutoMigrate(&entities.User{})
}
