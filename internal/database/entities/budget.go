package entities

import (
	"gorm.io/gorm"
)

type ETypeBudget string

const (
	Percent ETypeBudget = "percent"
	Amount  ETypeBudget = "amount"
)

type Budget struct {
	gorm.Model
	Name        string
	Description string
	Type        ETypeBudget
	Value       float32
	UserID      uint
	Active      bool
}
