package usecases

import (
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/vladjong/hand_card/config"
	postgresdb "github.com/vladjong/hand_card/internal/adapter/db/postgres_db"
	"github.com/vladjong/hand_card/internal/controller/http/v1/dto"
	"github.com/vladjong/hand_card/pkg/postgres"
)

func TestSignUP(t *testing.T) {
	cfg := config.GetConfig()
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading env variables: %s", err.Error())
	}
	postgresClient, _ := postgres.NewClient(
		postgres.PostgresConfig{
			Host:     cfg.PostgresSQL.Host,
			Port:     cfg.PostgresSQL.Port,
			Username: cfg.PostgresSQL.Username,
			Password: os.Getenv("POSTGRES_PASSWORD"),
			DBName:   cfg.PostgresSQL.DBName,
			SSLMode:  cfg.PostgresSQL.SSLMode,
		})
	storage := postgresdb.New(postgresClient)
	auth := NewAuthUseCase(storage)
	signUpDto := dto.SignUpDto{
		Login:    "test",
		Password: "1sd23",
		Email:    "test@mail.ru",
	}
	err := auth.CreateUser(signUpDto)
	assert.Nil(t, err)
}
