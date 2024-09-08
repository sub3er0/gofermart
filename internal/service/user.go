package service

import (
	"fmt"
	"gofermart/internal/repository"
)

type UserService struct {
	UserRepository *repository.UserRepository
}

func (us *UserService) IsUserExist(username string) bool {
	isExist, err := us.UserRepository.IsUserExists(username)
	if err != nil {
		return false
	}

	fmt.Println(isExist)
	return true
}
