package entities

type Card struct {
	Id           int    `json:"-" db:"id"`
	Organization string `json:"organization" binding:"required" db:"organization"`
	Name         string `json:"name" db:"name"`
}
