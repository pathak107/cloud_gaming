package service

// import (
// 	"errors"

// 	"github.com/pathak107/cloudesk/pkg/api/dto"
// 	"github.com/pathak107/cloudesk/pkg/api/entity"
// 	"gorm.io/gorm"
// )

// type AppsService struct {
// 	db *gorm.DB
// }

// func (s *AppsService) Create(user *dto.App) error {
// 	hash, err := s.authSvc.HashPassword(user.Password)
// 	if err != nil {
// 		return errors.New("password couldn't be generated")
// 	}
// 	res := s.db.Create(&entity.User{
// 		Name:     user.Name,
// 		Email:    user.Email,
// 		Password: hash,
// 		UserType: entity.UserType(user.UserType),
// 	})
// 	if res.Error != nil {
// 		return res.Error
// 	}
// 	return nil
// }

// func NewAppsService(db *gorm.DB) *AppsService {
// 	return &AppsService{
// 		db: db,
// 	}
// }
