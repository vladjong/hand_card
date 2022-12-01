package usecases

import (
	"crypto/sha1"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/vladjong/hand_card/internal/adapter/db"
	"github.com/vladjong/hand_card/internal/controller/http/v1/dto"
	"github.com/vladjong/hand_card/internal/entities"
)

const TokenTTL = 12 * time.Hour

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

func (a *authUseCase) GenerateToken(userDto dto.SingInDto) (token entities.Token, err error) {
	user := entities.User{
		Login:    userDto.Login,
		Password: generatePasswordHash(userDto.Password),
	}
	id, err := a.storage.GetUser(user)
	if err != nil {
		return token, err
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, &entities.TokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		UserId: id,
	})
	tokenName, err := tokenClaims.SignedString([]byte(os.Getenv("SIGNED_KEY")))
	if err != nil {
		return token, err
	}
	token.Name = tokenName
	return token, nil
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(os.Getenv("SALT"))))
}
