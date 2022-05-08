package entity

import "gorm.io/gorm"

type UserType string

const (
	Admin    UserType = "admin"
	Customer UserType = "customer"
	Editor   UserType = "editor"
)

type User struct {
	gorm.Model
	Name          *string        `gorm:"not null"`
	Email         *string        `gorm:"not null;unique"`
	Password      *string        `gorm:"not null"`
	UserType      UserType       `gorm:"not null;default:customer"`
	Subscriptions []Subscription `gorm:"many2many:subscriptions;"`
	Transactions  []Transaction
}
