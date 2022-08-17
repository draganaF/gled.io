package controller

import (
	"net/http"

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
