package db

import "github.com/vladjong/hand_card/internal/entities"

type Storager interface {
	CreateUser(user entities.User) error
}
