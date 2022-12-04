package postgresdb

import (
	"github.com/sirupsen/logrus"
	"github.com/vladjong/hand_card/internal/entities"
)

func (s *postgresStorage) CreateCard(card entities.Card, userId int) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}

	var categoryId int
	queryCategory := `INSERT INTO categories (category_name)
						VALUES ($1)
						ON CONFLICT(category_name)
						DO UPDATE SET category_name=EXCLUDED.category_name
						RETURNING id`
	row := tx.QueryRow(queryCategory, card.Category.CategoryName)
	if err := row.Scan(&categoryId); err != nil {
		if rb := tx.Rollback(); rb != nil {
			return rb
		}
		return err
	}
	logrus.Info(categoryId)

	var cardId int
	queryInsertCard := `INSERT INTO cards (number, organization, category_id)
						VALUES ($1, $2, $3)
						ON CONFLICT(number)
						DO UPDATE SET number=EXCLUDED.number
						RETURNING id`
	row = tx.QueryRow(queryInsertCard, card.Number, card.Organization, categoryId)
	if err := row.Scan(&cardId); err != nil {
		if rb := tx.Rollback(); rb != nil {
			return rb
		}
		return err
	}

	logrus.Info(cardId)

	queryInsertUserCard := `INSERT INTO user_cards (user_id, card_id)
							VALUES ($1, $2)
							ON CONFLICT(card_id)
							DO UPDATE SET card_id=EXCLUDED.card_id`
	if _, err := tx.Exec(queryInsertUserCard, userId, cardId); err != nil {
		if rb := tx.Rollback(); rb != nil {
			return rb
		}
		return err
	}
	return tx.Commit()
}

func (s *postgresStorage) GetCards(userId int) (cards []entities.Card, err error) {
	query := `SELECT c.number, c.organization, ct.category_name
				FROM user_cards AS uc
				JOIN cards c ON uc.card_id = c.id
				JOIN categories ct ON c.category_id = ct.id
				WHERE user_id=$1`
	if err := s.db.Select(&cards, query, userId); err != nil {
		return cards, err
	}
	return cards, nil
}
