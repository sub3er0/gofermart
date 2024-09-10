package repository

import (
	"gofermart/db"
	"gofermart/internal/models"
)

type UserRepository struct {
	DBStorage *db.PgStorage
}

func (ur *UserRepository) CreateUser(user models.User) error {
	query := "INSERT INTO users (username, password) VALUES ($1, $2)"
	_, err := ur.DBStorage.Conn.Exec(ur.DBStorage.Ctx, query, user.Username, user.Password)
	return err
}

func (ur *UserRepository) IsUserExists(username string) (bool, error) {
	var id int
	query := "SELECT id FROM users WHERE username = $1"
	err := ur.DBStorage.Conn.QueryRow(ur.DBStorage.Ctx, query, username).Scan(&id)

	if err != nil && err.Error() != "no rows in result set" {
		return false, err
	}

	if err != nil && err.Error() == "no rows in result set" {
		return false, nil
	}

	return true, nil
}

func (ur *UserRepository) GetUserByUsername(username string) (models.User, error) {
	var user models.User
	err := ur.DBStorage.Conn.QueryRow(
		ur.DBStorage.Ctx,
		"SELECT id, username, password FROM users WHERE username = $1", username).
		Scan(&user.ID, &user.Username, &user.Password)
	return user, err
}
