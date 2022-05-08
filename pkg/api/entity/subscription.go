package entity

import (
	"time"

	"gorm.io/gorm"
)

type Subscription struct {
	UserID           int `gorm:"primaryKey"`
	MembershipPlanID int `gorm:"primaryKey"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        gorm.DeletedAt `gorm:"index"`
	StartTime        time.Time
	ExpiryTime       time.Time
	TransactionID    int
	Transaction      Transaction
}
