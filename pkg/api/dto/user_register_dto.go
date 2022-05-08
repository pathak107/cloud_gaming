package dto

type User struct {
	Name     *string `json:"name" binding:"required"`
	Email    *string `json:"email" binding:"required,email"`
	Password *string `json:"password" binding:"required"`
	UserType string  `json:"user_type" binding:"required,oneof=admin customer editor"`
}
