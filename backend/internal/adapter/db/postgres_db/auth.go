package postgresdb

import (
	"github.com/vladjong/hand_card/internal/entities"
)

func (s *postgresStorage) CreateUser(user entities.User) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	queryInsert := `INSERT INTO users (email, password_hash, login)
						VALUES ($1, $2, $3)`
	if _, err := tx.Exec(queryInsert, user.Email, user.Password, user.Login); err != nil {
		if rb := tx.Rollback(); rb != nil {
			return rb
		}
		return err
	}
	return tx.Commit()
}

func (s *postgresStorage) GetUser(user entities.User) (id int, err error) {
	query := `SELECT id FROM users WHERE login=$1 AND password_hash=$2`
	if err := s.db.Get(&id, query, user.Login, user.Password); err != nil {
		return 0, err
	}
	return id, nil
}
