package dto

import "mime/multipart"

type CategoryUpdate struct {
	Name        *string               `form:"name"`
	Description *string               `form:"description"`
	Image       *multipart.FileHeader `form:"image"`
	ImageUrl    *string               `binding:"-"`
}
