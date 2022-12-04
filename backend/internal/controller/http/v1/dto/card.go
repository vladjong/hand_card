package dto

type CardDto struct {
	Organization string `json:"organization" binding:"required"`
	Name         string `json:"name"`
	CategoryName string `json:"category_name" binding:"required"`
}
