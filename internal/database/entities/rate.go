package entities

import (
	"gorm.io/gorm"
)

type Rate struct {
	gorm.Model
	UserID       uint
	Active       bool
	Value        float32
	InverseValue float64
	CurrencyID   uint
}
