package entity

import "gorm.io/gorm"

type MembershipPlan struct {
	gorm.Model
	Name            string `gorm:"not null"`
	Description     string `gorm:"not null"`
	Slug            string `gorm:"not null"`
	Cost            float64
	Duration        int //number of days
	HoursLimit      int
	VmsLimit        int
	AllowedAppPacks []AppsPack `gorm:"many2many:membership_appspack;"`
	Users           []User     `gorm:"many2many:subscriptions;"`
}
