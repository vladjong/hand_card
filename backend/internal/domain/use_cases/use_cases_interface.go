package usecases

import "github.com/vladjong/hand_card/internal/controller/http/v1/dto"

type AuthUseCaser interface {
	CreateUser(userDto dto.SignUpDto) error
}
