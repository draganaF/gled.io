package controller

import (
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	apicontract "github.com/draganaF/gled.io/UserService/apiContract"
	"github.com/draganaF/gled.io/UserService/service"
	"github.com/draganaF/gled.io/UserService/utils"
)

var Auth = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	var loginRequest *apicontract.Login

	err := utils.ReadJSONBody(r, &loginRequest)

	if err != nil {
		utils.JSONResponse(w, 400, "Bad request.")
	}

	authService := service.NewAuthService()

	jwtToken, err := authService.Auth(loginRequest.Email, loginRequest.Password)

	if err != nil {
		utils.JSONResponse(w, 401, "Not authorized")
		return
	}

	utils.JSONResponse(w, 200, jwtToken)
})

var AuthorizeUser = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	tokenString := r.Header.Get("Authorization")
	if len(tokenString) == 0 {
		utils.JSONResponse(w, 400, "No header")
		return
	}

	tokenString = strings.Replace(tokenString, "Bearer ", "", 1)
	claims, err := utils.VerifyJwtToken(tokenString)
	if err != nil {
		utils.JSONResponse(w, 401, "You are not authorized")
		return
	}

	utils.JSONResponse(w, 200, &apicontract.AuthorizedUser{
		Name: claims.(jwt.MapClaims)["name"].(string),
		Role: claims.(jwt.MapClaims)["role"].(string),
	})
})
