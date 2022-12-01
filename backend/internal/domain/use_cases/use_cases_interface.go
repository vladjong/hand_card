package usecases

import (
	"github.com/vladjong/hand_card/internal/controller/http/v1/dto"
	"github.com/vladjong/hand_card/internal/entities"
)

type AuthUseCaser interface {
	CreateUser(userDto dto.SignUpDto) error
	GenerateToken(userDto dto.SingInDto) (entities.Token, error)
}