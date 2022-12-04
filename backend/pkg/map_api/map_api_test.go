package mapapi

import (
	"log"
	"os"
	"sync"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	gisapi "github.com/vladjong/hand_card/pkg/map_api/gis_api"
)

const (
	lat = 82.897918
	lon = 54.980332
)

func TestWithCoordinate(t *testing.T) {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading env variables: %s", err.Error())
	}
	cfg := gisapi.Config{
		Key:    os.Getenv("GIS_KEY"),
		Radius: 250,
		Sort:   "distance"}
	map_api := gisapi.New(cfg)
	categories := map[string]bool{"Еда": true}
	organisations, err := map_api.GetCompany(categories, lat, lon)
	assert.NotEmpty(t, organisations["Еда"])
	assert.Nil(t, err)
}

func TestWithIncorectCategory(t *testing.T) {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading env variables: %s", err.Error())
	}
	cfg := gisapi.Config{
		Key:    os.Getenv("GIS_KEY"),
		Radius: 250,
		Sort:   "distance"}
	map_api := gisapi.New(cfg)
	categories := map[string]bool{"ФЫВцй": true}
	organisations, err := map_api.GetCompany(categories, lat, lon)
	assert.Empty(t, organisations["ФЫВцй"])
	assert.Nil(t, err)
}

func BenchmarWithCoordinate(b *testing.B) {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading env variables: %s", err.Error())
	}
	cfg := gisapi.Config{
		Key:    os.Getenv("GIS_KEY"),
		Radius: 250,
		Sort:   "distance"}
	map_api := gisapi.New(cfg)
	categories := map[string]bool{"Еда": true, "Спорт": true, "Магазин": true, "Пицца": true}
	for i := 0; i < b.N; i++ {
		wg := sync.WaitGroup{}
		for i := 0; i < 100; i++ {
			wg.Add(1)
			go func() {
				_, err := map_api.GetCompany(categories, lat, lon)
				if err != nil {
					panic(err)
				}
				wg.Done()
			}()
		}
		wg.Wait()
	}
}
