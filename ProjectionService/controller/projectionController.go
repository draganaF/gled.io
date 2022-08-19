package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	apicontract "github.com/draganaF/gled.io/ProjectionService/apiContract"
	"github.com/draganaF/gled.io/ProjectionService/service"
	"github.com/draganaF/gled.io/ProjectionService/utils"
	"github.com/gorilla/mux"
)

var SearchProjections = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	var p = apicontract.SearchParams{}

	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	projectionService := service.NewProjectionService()

	projections, err := projectionService.Search(&p)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	utils.JSONResponse(w, 200, projections)
})

var ReadAllProjections = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	projectionService := service.NewProjectionService()

	projections, err := projectionService.ReadAll()

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	utils.JSONResponse(w, 200, projections)
})

var ReadProjectionById = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	projectionService := service.NewProjectionService()

	projection, err := projectionService.ReadProjectionById(uint(id))

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	utils.JSONResponse(w, 200, projection)
})

var ReadProjectionByMovieId = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	projectionService := service.NewProjectionService()

	projections, err := projectionService.ReadProjectionsByMovieId(uint(id))

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	utils.JSONResponse(w, 200, projections)
})

var CreateProjection = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	var request *apicontract.CreateProjectionRequest
	err := utils.ReadJSONBody(r, &request)

	if err != nil {
		return
	}

	projectionService := service.NewProjectionService()

	createdProjection, err := projectionService.Create(*request)

	if err != nil {
		utils.JSONResponse(w, 400, "Bad request")
		return
	}

	utils.JSONResponse(w, 200, createdProjection)
})

var UpdateProjection = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	var request *apicontract.UpdateProjectionRequest
	err := utils.ReadJSONBody(r, &request)

	if err != nil {
		return
	}

	projectionService := service.NewProjectionService()

	updatedProjection, err := projectionService.Update(*request)

	if err != nil {
		utils.JSONResponse(w, 400, "Bad request")
		return
	}

	utils.JSONResponse(w, 200, updatedProjection)
})

var DeleteProjection = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	projectionService := service.NewProjectionService()

	projectionService.Delete(uint(id))

	utils.JSONResponse(w, 200, nil)
})
