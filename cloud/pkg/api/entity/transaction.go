package entity

import (
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	Amount float64
	Token  string
	UserID int
	User   User
	Info   string
}
