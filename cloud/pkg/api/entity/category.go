package entity

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name        *string `gorm:"not null;unique"`
	Description *string `gorm:"not null" `
	Slug        *string `gorm:"not null;unique" `
	ImageUrl    *string
	Apps        []App
}
