package controller

import (
	"fmt"
	"net/http"
	"strconv"

	apicontract "github.com/draganaF/gled.io/ProjectionService/apiContract"
	"github.com/draganaF/gled.io/ProjectionService/service"
	"github.com/draganaF/gled.io/ProjectionService/utils"
	"github.com/gorilla/mux"
)

var ReadAllMovies = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	movieService := service.NewMovieService()

	movies, err := movieService.ReadAll()

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	utils.JSONResponse(w, 200, movies)
})

var ReadMovieById = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	movieService := service.NewMovieService()

	movie, err := movieService.ReadMovieById(uint(id))

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	utils.JSONResponse(w, 200, movie)
})

var CreateMovie = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	var request *apicontract.CreateMovieRequest
	fmt.Println("E JEBO JA SEBE")
	err := utils.ReadJSONBody(r, &request)

	if err != nil {
		return
	}

	movieService := service.NewMovieService()

	createdMovie, err := movieService.Create(*request)

	fmt.Println("U kontroleru sam")
	fmt.Println(createdMovie)

	if err != nil {
		utils.JSONResponse(w, 400, "Bad request")
		return
	}

	utils.JSONResponse(w, 200, createdMovie)
})

var UpdateMovie = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	fmt.Println("UPDATE")
	var request *apicontract.UpdateMovieRequest
	err := utils.ReadJSONBody(r, &request)

	if err != nil {
		return
	}

	movieService := service.NewMovieService()

	updatedMovie, err := movieService.Update(*request)

	if err != nil {
		utils.JSONResponse(w, 400, "Bad request")
		return
	}

	utils.JSONResponse(w, 200, updatedMovie)
})

var DeleteMovie = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	movieService := service.NewMovieService()

	movieService.Delete(uint(id))

	utils.JSONResponse(w, 200, nil)
})
