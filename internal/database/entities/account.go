package entities

import (
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	Name    string
	Balance float64
	Icon    string
	Active  bool
	UserID  uint
}
