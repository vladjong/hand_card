package entities

type Card struct {
	Id           int    `json:"-" db:"id"`
	Organization string `json:"organization" binding:"required" db:"organization"`
	Name         string `json:"name" db:"name"`
	Category
}

type Category struct {
	Id           int    `json:"-" db:"id"`
	CategoryName string `json:"category_name" db:"category_name"`
}
