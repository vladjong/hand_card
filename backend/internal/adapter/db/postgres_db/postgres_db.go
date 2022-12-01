package postgresdb

import "github.com/jmoiron/sqlx"

type postgresStorage struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *postgresStorage {
	return &postgresStorage{db}
}
