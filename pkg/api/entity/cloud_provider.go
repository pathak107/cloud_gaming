package entity

import (
	"github.com/pathak107/cloudesk/pkg/cloud"
	"gorm.io/gorm"
)

type CloudProvider struct {
	gorm.Model
	Name    string `gorm:"not null"`
	CloudID cloud.CloudType
}
