package database

import (
	"github.com/pathak107/cloudesk/pkg/api/entity"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DB struct {
	Conn *gorm.DB
}

func NewConnection() (*DB, error) {
	conn, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	//Migrate the schema
	err = conn.AutoMigrate(
		&entity.AppsPack{},
		&entity.User{},
		&entity.Category{},
		&entity.App{},
		&entity.AppsPackCost{},
		&entity.CloudProvider{},
		&entity.MembershipPlan{},
		&entity.VM{},
		&entity.Transaction{},
		&entity.Subscription{},
		&entity.RDP{},
	)
	if err != nil {
		return nil, err
	}

	// err = conn.SetupJoinTable(, "Addresses", &PersonAddress{})

	return &DB{
		Conn: conn,
	}, nil
}

// func (db *DB) GetModel() *gorm.Model {
// 	return
// }
