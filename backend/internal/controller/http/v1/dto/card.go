package dto

type CardDto struct {
	Organization string `json:"organization" binding:"required"`
	Name         string `json:"name"`
}
