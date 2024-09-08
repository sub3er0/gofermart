package repository

import (
	"database/sql"
	"errors"
	"gofermart/db"
)

type UserRepository struct {
	DBStorage *db.PgStorage
}

func (ur *UserRepository) CreateUser() error {
	query := "INSERT INTO users (username, password, email) VALUES ($1, $2, $3)"
	_, err := ur.DBStorage.Exec(query)
	return err
}

func (ur *UserRepository) IsUserExists(username string) (bool, error) {
	var id int
	query := "SELECT id FROM users WHERE username = $1"
	err := ur.DBStorage.Conn.QueryRow(ur.DBStorage.Ctx, query, username).Scan(&id)

	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return false, err
	}

	if errors.Is(err, sql.ErrNoRows) {
		return false, nil
	}

	return true, nil
}
