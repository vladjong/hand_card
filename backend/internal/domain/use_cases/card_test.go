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
	gisapi "github.com/vladjong/hand_card/pkg/map_api/gis_api"
	"github.com/vladjong/hand_card/pkg/postgres"
)

func TestPostCard(t *testing.T) {
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
	gisApi := gisapi.New(gisapi.Config{
		Key:    os.Getenv("GIS_KEY"),
		Radius: 250,
		Sort:   "distance",
	})
	card := NewCardUseCase(storage, gisApi)
	cardDto := dto.CardDto{
		Organization: "test",
		Number:       "1sd23",
		CategoryName: "Спорт",
	}
	err := card.CreateCard(cardDto, 1)
	assert.Nil(t, err)
}

func TestGetCard(t *testing.T) {
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
	gisApi := gisapi.New(gisapi.Config{
		Key:    os.Getenv("GIS_KEY"),
		Radius: 250,
		Sort:   "distance",
	})
	card := NewCardUseCase(storage, gisApi)
	_, err := card.GetCards(1, dto.Coordinate{})
	assert.Nil(t, err)
}

func TestDeleteCard(t *testing.T) {
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
	gisApi := gisapi.New(gisapi.Config{
		Key:    os.Getenv("GIS_KEY"),
		Radius: 250,
		Sort:   "distance",
	})
	card := NewCardUseCase(storage, gisApi)
	err := card.DeleteCard(1, 1)
	assert.Nil(t, err)
}
