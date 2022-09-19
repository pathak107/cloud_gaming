package entity

import "gorm.io/gorm"

type AppsPack struct {
	gorm.Model
	Name        string `gorm:"not null"`
	Description string `gorm:"not null"`
	Slug        string `gorm:"not null"`
	ConfigJson  string
	Apps        []App
	RDPs        []RDP `gorm:"many2many:appspack_rdps;"`
}
