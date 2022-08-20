package controller

import (
	"net/http"
	"strconv"

	"github.com/draganaF/gled.io/ProjectionService/service"
	"github.com/draganaF/gled.io/ProjectionService/utils"
	"github.com/gorilla/mux"
)

var GetAllCinemaHalls = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	hallsService := service.NewCinemaHallService()

	halls, err := hallsService.ReadCinemaHallsAll()

	if err != nil {
		utils.JSONResponse(w, 400, "Bad request")
		return
	}
	utils.JSONResponse(w, 200, halls)

})

var GetCinemaHallById = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	hallsService := service.NewCinemaHallService()

	hall, err := hallsService.ReadCinemaHallById(uint(id))

	if err != nil {
		utils.JSONResponse(w, 400, "Bad request")
		return
	}

	utils.JSONResponse(w, 200, hall)
})
