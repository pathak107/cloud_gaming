package dto

type CreateVmDTO struct {
	Name     string `json:"name" binding:"required"`
	Image    string `json:"image" binding:"required"`
	Hardware string `json:"hardware" binding:"required"`
	Storage  string `json:"storage" binding:"required"`
}
