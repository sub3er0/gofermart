package service

import (
	"gofermart/internal/models"
	"gofermart/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	UserRepository *repository.UserRepository
}

func (us *UserService) IsUserExist(username string) int {
	return us.UserRepository.IsUserExists(username)
}

func (us *UserService) RegisterUser(user models.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)

	return us.UserRepository.CreateUser(user)
}

func (us *UserService) AuthenticateUser(username, password string) (models.User, error) {
	user, err := us.UserRepository.GetUserByUsername(username)

	if err != nil {
		return models.User{}, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return models.User{}, err
	}

	return user, nil
}
