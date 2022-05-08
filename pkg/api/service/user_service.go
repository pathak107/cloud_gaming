package service

import (
	"github.com/pathak107/cloudesk/pkg/api/apierrors"
	"github.com/pathak107/cloudesk/pkg/api/dto"
	"github.com/pathak107/cloudesk/pkg/api/entity"
	"github.com/pathak107/cloudesk/pkg/api/helpers"
	"gorm.io/gorm"
)

type UserService struct {
	db      *gorm.DB
	authSvc *AuthService
}

func (s *UserService) Create(user *dto.User) error {
	hash, err := s.authSvc.HashPassword(helpers.ToString(user.Password))
	if err != nil {
		return err
	}
	res := s.db.Create(&entity.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: helpers.StringPtr(hash),
		UserType: entity.UserType(user.UserType),
	})
	if res.Error != nil {
		return apierrors.NewServerError(res.Error, "user_creation")
	}
	return nil
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{
		db: db,
	}
}
