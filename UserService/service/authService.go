package service

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/draganaF/gled.io/UserService/repository"
	"github.com/draganaF/gled.io/UserService/utils"
)

type AuthService struct {
	repository *repository.UserRepository
}

func NewAuthService() *AuthService {
	return &AuthService{
		repository: &repository.UserRepository{
			DB: utils.DB,
		},
	}
}

func (authService *AuthService) Auth(email string, password string) (string, error) {
	user := authService.repository.ReadUserByEmail(email)

	if user == nil {
		return "", errors.New("User does not exist.")
	}

	if user.Password != password {
		return "", errors.New("Passwords do not match.")
	}
	fmt.Println("Rola u useru je " + strconv.Itoa(int(user.Role)))
	claims := utils.Claims{
		Name: strconv.FormatUint(uint64(user.Id), 10),
		Role: strconv.Itoa(int(user.Role)),
	}
	jwtToken, _ := utils.GetJwtToken(claims)

	return jwtToken, nil
}
