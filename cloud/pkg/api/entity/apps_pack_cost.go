package entity

import (
	"time"

	"gorm.io/gorm"
)

type AppsPackCost struct {
	AppsPackID      int     `gorm:"primaryKey"`
	CloudProviderID int     `gorm:"primaryKey"`
	Cost            float64 //hourly cost
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt `gorm:"index"`
}
