package handlers

import (
	"encoding/json"
	"fmt"
	"gofermart/internal/service"
	"net/http"
)

type UserHandler struct {
	UserService service.UserService
}

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (us *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userExist, err := us.UserService.UserRepository.IsUserExists(user.Username)

	if err != nil {

	}

	fmt.Println(userExist)

}

func (us *UserHandler) Login(w http.ResponseWriter, r *http.Request) {

}
