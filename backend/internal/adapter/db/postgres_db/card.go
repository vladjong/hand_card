package postgresdb

import "github.com/vladjong/hand_card/internal/entities"

func (s *postgresStorage) CreateCard(card entities.Card, userId int) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	var cardId int
	queryInsertCard := `INSERT INTO cards (name, organization)
						VALUES ($1, $2) RETURNING id`
	row := tx.QueryRow(queryInsertCard, card.Name, card.Organization)
	if err := row.Scan(&cardId); err != nil {
		if rb := tx.Rollback(); rb != nil {
			return rb
		}
		return err
	}
	queryInsertUserCard := `INSERT INTO user_cards (user_id, card_id)
							VALUES ($1, $2)`
	if _, err := tx.Exec(queryInsertUserCard, userId, cardId); err != nil {
		if rb := tx.Rollback(); rb != nil {
			return rb
		}
		return err
	}
	return tx.Commit()
}

func (s *postgresStorage) GetCards(userId int) (cards []entities.Card, err error) {
	query := `SELECT c.name, c.organization
				FROM user_cards AS uc
				JOIN cards c ON uc.card_id = c.id
				WHERE user_id=$1`
	if err := s.db.Select(&cards, query, userId); err != nil {
		return cards, err
	}
	return cards, nil
}
