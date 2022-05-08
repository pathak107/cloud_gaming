package entity

import "gorm.io/gorm"

type RDP struct {
	gorm.Model
	Name        string `gorm:"not null"`
	Description string `gorm:"not null"`
	Slug        string `gorm:"not null"`
	ImageUrl    string
}
