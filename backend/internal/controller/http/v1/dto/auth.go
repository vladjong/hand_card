package dto

type SignUpDto struct {
	Login    string `json:"login"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
}

type SingInDto struct {
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type TokenDto struct {
	Token string `json:"token"`
}
