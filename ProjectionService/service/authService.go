package service

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	apicontract "github.com/draganaF/gled.io/ProjectionService/apiContract"
)

type AuthService struct {
	client *http.Client
}

func NewAuthService() *AuthService {
	return &AuthService{
		client: &http.Client{},
	}
}

func (authService *AuthService) Authorize(token string, roles *[]int) (*apicontract.AuthorizedUser, error) {
	url := os.Getenv("USER_SERVICE_URL") + "/api/authorize"

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", token)
	req.Header.Add("Accept", "application/json")

	res, err := authService.client.Do(req)
	if err != nil {
		return nil, errors.New("you are not authorized")
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, errors.New("could not authorize")
	}

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	user := &apicontract.AuthorizedUser{}
	json.Unmarshal(bodyBytes, user)

	value, _ := strconv.Atoi(user.Role)
	if authService.checkRoles(value, roles) {
		return user, nil
	}

	return nil, errors.New("you don't have permission")
}

func (authService *AuthService) checkRoles(role int, roles *[]int) bool {

	for _, r := range *roles {
		if r == role {
			return true
		}
	}

	return false
}
