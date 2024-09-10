package entities

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username   string
	Password   string
	Active     bool
	Budgets    []Budget    `gorm:"foreignKey:UserID"`
	Rates      []Rate      `gorm:"foreignKey:UserID"`
	Operations []Operation `gorm:"foreignKey:UserID"`
}
