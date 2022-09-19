package service

import (
	"errors"

	"github.com/gosimple/slug"
	"github.com/pathak107/cloudesk/pkg/api/apierrors"
	"github.com/pathak107/cloudesk/pkg/api/dto"
	"github.com/pathak107/cloudesk/pkg/api/entity"
	"github.com/pathak107/cloudesk/pkg/api/helpers"
	"gorm.io/gorm"
)

type CategoryService struct {
	db *gorm.DB
}

func (s *CategoryService) Create(category *dto.CategoryCreate) error {
	res := s.db.Create(&entity.Category{
		Name:        category.Name,
		Description: category.Description,
		Slug:        helpers.StringPtr(slug.Make(helpers.ToString(category.Name))),
		ImageUrl:    category.ImageUrl,
	})
	if res.Error != nil {
		return apierrors.NewServerError(res.Error, "category_creation")
	}
	return nil
}

func (s *CategoryService) FindOne(catID string) (entity.Category, error) {
	var category entity.Category
	res := s.db.First(&category, catID)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return category, apierrors.NewResourceNotFoundError(res.Error, "category_find_one", "category does not exist")
		}
		return category, apierrors.NewServerError(res.Error, "category_find_one")
	}
	return category, nil
}

func (s *CategoryService) FindAll() ([]entity.Category, error) {
	var category []entity.Category
	res := s.db.Find(&category)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return category, apierrors.NewResourceNotFoundError(res.Error, "category_find", "no categories to show")
		}
		return category, apierrors.NewServerError(res.Error, "category_find")
	}
	return category, nil
}

func (s *CategoryService) Update(catUpdate *dto.CategoryUpdate, catID string) error {
	category, err := s.FindOne(catID)
	if err != nil {
		return err
	}
	category.Name = catUpdate.Name
	category.Description = catUpdate.Description
	category.ImageUrl = catUpdate.ImageUrl
	category.Slug = helpers.StringPtr(slug.Make(helpers.ToString(catUpdate.Name)))
	res := s.db.Save(&category)
	if res.Error != nil {
		return apierrors.NewServerError(res.Error, "category_update")
	}
	return nil
}

func (s *CategoryService) Delete(catID string) error {
	if _, err := s.FindOne(catID); err != nil {
		return err
	}
	res := s.db.Delete(&entity.Category{}, catID)
	if res.Error != nil {
		return apierrors.NewServerError(res.Error, "category_delete")
	}
	return nil
}

func NewCategoryService(db *gorm.DB) *CategoryService {
	return &CategoryService{
		db: db,
	}
}
