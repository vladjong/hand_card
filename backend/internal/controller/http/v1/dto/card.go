package dto

type CardDto struct {
	Id           int    `json:"id"`
	Organization string `json:"organization" binding:"required"`
	Number       string `json:"number"`
	CategoryName string `json:"category_name" binding:"required"`
}
