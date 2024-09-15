package entities

import (
	"gorm.io/gorm"
)

type ETypeOperations string

const (
	Income   ETypeOperations = "income"
	Outgoing ETypeOperations = "outgoing"
	Transfer ETypeOperations = "transfer"
)

type Operation struct {
	gorm.Model
	UserID uint
	Active bool
	Type   ETypeOperations
}
