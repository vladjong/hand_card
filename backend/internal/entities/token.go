package entities

import "github.com/golang-jwt/jwt"

type Token struct {
	Name string `json:"name"`
}

type TokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}
