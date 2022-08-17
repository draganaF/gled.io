package router

import (
	"log"
	"net/http"

	"github.com/draganaF/gled.io/ProjectionService/controller"
	"github.com/gorilla/mux"
)

func HandleRequests() {
	router := mux.NewRouter()

	router.Handle("/api/movies", controller.ReadAllMovies).Methods("GET")
	router.Handle("/api/movies/{id}", controller.ReadMovieById).Methods("GET")
	router.Handle("/api/movies", controller.CreateMovie).Methods("POST")
	router.Handle("/api/movies", controller.UpdateMovie).Methods("PUT")
	router.Handle("/api/movies/{id}", controller.DeleteMovie).Methods("DELETE")

	router.Handle("/api/projections", controller.ReadAllProjections).Methods("GET")
	router.Handle("/api/projections/{id}", controller.ReadProjectionById).Methods("GET")
	router.Handle("/api/projections", controller.CreateProjection).Methods("POST")
	router.Handle("/api/projections", controller.UpdateProjection).Methods("PUT")
	router.Handle("/api/projections/{id}", controller.DeleteProjection).Methods("DELETE")

	router.Handle("/api/projections/search", controller.SearchProjections).Methods("GET")

	// corsOpts := cors.New(cors.Options{
	// 	AllowedOrigins: []string{"*"},
	// 	AllowedHeaders: []string{"*"},
	// 	AllowedMethods: []string{
	// 		http.MethodGet,
	// 		http.MethodPost,
	// 		http.MethodPut,
	// 		http.MethodPatch,
	// 		http.MethodDelete,
	// 		http.MethodOptions,
	// 		http.MethodHead,
	// 	},
	// })
	log.Fatal(http.ListenAndServe(":8082", router))
}
