package services

import (
	models "finances/internal/models"
	entities "finances/internal/database/entities"
	database "finances/internal/database"
	libs "finances/pkg/lib"
)

func CreateUserService(user *models.MUser) (entities.User, error) {
	pass, err := libs.HashString(user.Password)

	if err != nil {
		return entities.User{}, err
	}

	newUser := entities.User{Username: user.Username, Password: pass}

	db := database.DB

	result := db.Create(&newUser)

	if result.Error != nil {
		return entities.User{}, result.Error
	}

	return newUser, nil
}