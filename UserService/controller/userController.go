package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	apicontract "github.com/draganaF/gled.io/UserService/apiContract"
	"github.com/draganaF/gled.io/UserService/model"
	"github.com/draganaF/gled.io/UserService/service"
	"github.com/draganaF/gled.io/UserService/utils"
	"github.com/gorilla/mux"
)

var Search = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	var p = apicontract.SearchParams{}

	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userService := service.NewUserService()

	users, err := userService.Search(&p)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	utils.JSONResponse(w, 200, users)
})

var ReadUserById = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	userService := service.NewUserService()

	user, err := userService.ReadUserById(uint(id))

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}

	utils.JSONResponse(w, 200, user)
})

var ReadUserByEmail = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	email, _ := params["email"]

	userService := service.NewUserService()

	user, err := userService.ReadUserByEmail(email)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}

	utils.JSONResponse(w, 200, user)
})

var CreateUser = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	var request *apicontract.CreateUser
	err := utils.ReadJSONBody(r, &request)

	if err != nil {
		return
	}

	userService := service.NewUserService()
	activationService := service.NewActivationService()

	createdUser, err := userService.Create(*request)

	if err != nil {
		utils.JSONResponse(w, 400, "Bad request")
	}

	if createdUser.Role == model.RegisteredUser {
		activationService.GenerateActivationLink(createdUser.Id)
	}

	utils.JSONResponse(w, 200, createdUser)
})

var UpdateUser = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	var request *apicontract.UpdateUser
	err := utils.ReadJSONBody(r, &request)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	userService := service.NewUserService()

	updatedUser, err := userService.Update(*request)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	utils.JSONResponse(w, 200, updatedUser)
})

var ChangePassword = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	var request *apicontract.ChangePassword
	err := utils.ReadJSONBody(r, &request)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	userService := service.NewUserService()

	updatedUser, err := userService.ChangeUserPassword(*request)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	utils.JSONResponse(w, 200, updatedUser)
})

var IncrementNegativePoints = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	userService := service.NewUserService()

	updatedUser, err := userService.IncrementNegativePoints(uint(id))

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	utils.JSONResponse(w, 200, updatedUser)
})

var IncrementNumberOfBougthTickets = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	userService := service.NewUserService()

	updatedUser, err := userService.IncrementNumberOfBoughtTickets(uint(id))

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	utils.JSONResponse(w, 200, updatedUser)
})

var IncrementNumberOfReservedTickets = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	userService := service.NewUserService()

	updatedUser, err := userService.IncrementNumberOfReservedTickets(uint(id))

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	utils.JSONResponse(w, 200, updatedUser)
})

var IncrementNumberOfSoldTickets = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	userService := service.NewUserService()

	updatedUser, err := userService.IncrementNumberOfSoldTickets(uint(id))

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	utils.JSONResponse(w, 200, updatedUser)
})

var BlockUser = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	userService := service.NewUserService()

	blockedUser, err := userService.BlockUser(uint(id))

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	utils.JSONResponse(w, 200, blockedUser)
})

var DeleteUser = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	userService := service.NewUserService()

	userService.Delete(uint(id))

	utils.JSONResponse(w, 200, nil)
})

var ActivateUser = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	text := params["link"]

	userService := service.NewUserService()
	activationService := service.NewActivationService()

	link, err := activationService.ActivateUser(text)

	if err != nil {
		utils.JSONResponse(w, 400, "Bad Reaquest")
	}

	user, err := userService.ActivateUser(link.UserId)

	if err != nil {
		utils.JSONResponse(w, 400, "Bad Reaquest")
	}

	utils.JSONResponse(w, 200, user)
})
