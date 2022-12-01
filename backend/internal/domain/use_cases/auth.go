package usecases

import (
	"crypto/sha1"
	"fmt"
	"os"

	"github.com/vladjong/hand_card/internal/adapter/db"
	"github.com/vladjong/hand_card/internal/controller/http/v1/dto"
	"github.com/vladjong/hand_card/internal/entities"
)

type authUseCase struct {
	storage db.Storager
}

func NewAuthUseCase(storage db.Storager) *authUseCase {
	return &authUseCase{storage}
}

func (a *authUseCase) CreateUser(userDto dto.SignUpDto) error {

	user := entities.User{
		Login:    userDto.Login,
		Password: generatePasswordHash(userDto.Password),
		Email:    userDto.Email,
	}
	if err := a.storage.CreateUser(user); err != nil {
		return err
	}
	return nil
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(os.Getenv("SALT"))))
}
