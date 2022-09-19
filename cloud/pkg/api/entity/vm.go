package entity

import (
	"time"

	"gorm.io/gorm"
)

type VMState int

const (
	Running VMState = iota
	Stopped
	Deleted
	Initializing
)

type VM struct {
	gorm.Model
	VmID        string  `gorm:"not null;unique"`
	Name        string  `gorm:"not null"`
	PublicIP    string  `gorm:"not null;unique"`
	Username    string  `gorm:"not null;"`
	Password    string  `gorm:"not null;"`
	GuacToken   string  `gorm:"not null;"`
	State       VMState `gorm:"default:3;"`
	Duration    int     // in seconds
	StartTime   time.Time
	StopTime    time.Time
	TimeElapsed int //in seconds
}
