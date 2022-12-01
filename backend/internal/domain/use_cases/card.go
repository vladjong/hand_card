package usecases

import (
	"github.com/vladjong/hand_card/internal/adapter/db"
	"github.com/vladjong/hand_card/internal/controller/http/v1/dto"
	"github.com/vladjong/hand_card/internal/entities"
)

type cardUseCase struct {
	storage db.Storager
}

func NewCardUseCase(storage db.Storager) *cardUseCase {
	return &cardUseCase{storage}
}

func (c *cardUseCase) CreateCard(cardDto dto.CardDto, userId int) error {
	card := entities.Card{
		Organization: cardDto.Organization,
		Name:         cardDto.Name,
	}
	if len(card.Name) == 0 {
		card.Name = cardDto.Organization
	}
	if err := c.storage.CreateCard(card, userId); err != nil {
		return err
	}
	return nil
}

func (c *cardUseCase) GetCards(userId int) (cards []entities.Card, err error) {
	cards, err = c.storage.GetCards(userId)
	return cards, err
}
