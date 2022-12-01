package postgresdb

import "github.com/vladjong/hand_card/internal/entities"

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
