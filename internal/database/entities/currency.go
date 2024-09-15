package entities

import (
	"gorm.io/gorm"
)

type Currency struct {
	gorm.Model
	Name   string
	Code   string
	Symbol string
	Active bool
	Rates  []Rate `gorm:"foreignKey:CurrencyID"`
}
