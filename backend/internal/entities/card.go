package entities

type Card struct {
	Id           int    `json:"-" db:"id"`
	Organization string `json:"organization" binding:"required" db:"organization"`
	Number       string `json:"number" db:"number"`
	Category
}

type Category struct {
	Id           int    `json:"-" db:"id"`
	CategoryName string `json:"category_name" db:"category_name"`
}
