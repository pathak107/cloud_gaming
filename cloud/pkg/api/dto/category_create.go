package dto

import "mime/multipart"

type CategoryCreate struct {
	Name        *string               `form:"name" binding:"required"`
	Description *string               `form:"description" binding:"required"`
	Image       *multipart.FileHeader `form:"image"`
	ImageUrl    *string               `binding:"-"`
}
