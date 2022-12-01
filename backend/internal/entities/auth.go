package entities

type User struct {
	Id       int    `json:"-" db:"id"`
	Login    string `json:"login" db:"login"`
	Password string `json:"password" binding:"required" db:"password_hash"`
	Email    string `json:"email" binding:"required" db:"email"`
}

type Token struct {
	Token string `json:"token"`
}
