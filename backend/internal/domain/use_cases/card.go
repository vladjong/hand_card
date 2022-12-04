package usecases

import (
	"fmt"
	"regexp"

	"github.com/vladjong/hand_card/internal/adapter/db"
	"github.com/vladjong/hand_card/internal/controller/http/v1/dto"
	"github.com/vladjong/hand_card/internal/entities"
	mapapi "github.com/vladjong/hand_card/pkg/map_api"
)

type cardUseCase struct {
	storage db.Storager
	mapApi  mapapi.IMapApi
}

func NewCardUseCase(storage db.Storager, mapApi mapapi.IMapApi) *cardUseCase {
	return &cardUseCase{
		storage: storage,
		mapApi:  mapApi,
	}
}

func (c *cardUseCase) CreateCard(cardDto dto.CardDto, userId int) error {
	card := entities.Card{
		Organization: cardDto.Organization,
		Number:       cardDto.Number,
		Category:     entities.Category{CategoryName: cardDto.CategoryName},
	}
	if len(card.Number) == 0 {
		card.Number = cardDto.Organization
	}
	if err := c.storage.CreateCard(card, userId); err != nil {
		return err
	}
	return nil
}

func (c *cardUseCase) GetCards(userId int, coordinate dto.Coordinate) ([]dto.CardDto, error) {
	cards, err := c.storage.GetCards(userId)
	if err != nil {
		return nil, err
	}
	if coordinate.Lat == 0 {
		return c.cardsToCardsDto(cards), nil
	}
	categories := make(map[string]bool)
	for _, card := range cards {
		categories[card.Category.CategoryName] = true
	}
	categoryCompanies, err := c.mapApi.GetCompany(categories, coordinate.Lat, coordinate.Lon)
	if err != nil {
		return nil, err
	}
	var mainCards []dto.CardDto

	for category, organisations := range categoryCompanies {
		for i, card := range cards {
			if category != card.Category.CategoryName {
				continue
			}
			cardDto, ok := c.find(organisations, card)
			if !ok {
				continue
			}
			deleteCard(cards, i)
			mainCards = append(mainCards, cardDto)
		}
	}
	return append(mainCards, c.cardsToCardsDto(cards)...), err
}

func (c *cardUseCase) cardsToCardsDto(cards []entities.Card) (cardsDto []dto.CardDto) {
	for _, card := range cards {
		if card.Organization == "" {
			continue
		}
		cardDto := dto.CardDto{
			Organization: card.Organization,
			Number:       card.Number,
			CategoryName: card.Category.CategoryName,
		}
		cardsDto = append(cardsDto, cardDto)
	}
	return cardsDto
}

func (c *cardUseCase) find(organisations []string, card entities.Card) (dto.CardDto, bool) {
	for _, organisation := range organisations {
		if ok, _ := regexp.MatchString(fmt.Sprintf("(?i)%s", card.Organization), organisation); ok {
			return dto.CardDto{
				Organization: card.Organization,
				Number:       card.Number,
				CategoryName: card.Category.CategoryName,
			}, true
		}
	}
	return dto.CardDto{}, false
}

func deleteCard(cards []entities.Card, i int) []entities.Card {
	cards[i] = cards[len(cards)-1]
	cards[len(cards)-1] = entities.Card{}
	cards = cards[:len(cards)-1]
	return cards
}
