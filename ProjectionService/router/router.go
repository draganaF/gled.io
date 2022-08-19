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
	router.Handle("/api/movies", Authenticate(controller.CreateMovie, 1)).Methods("POST")
	router.Handle("/api/movies", Authenticate(controller.UpdateMovie, 1)).Methods("PUT")
	router.Handle("/api/movies/{id}", Authenticate(controller.DeleteMovie, 1)).Methods("DELETE")

	router.Handle("/api/projections", controller.ReadAllProjections).Methods("GET")
	router.Handle("/api/projections/{id}", controller.ReadProjectionById).Methods("GET")
	router.Handle("/api/projections", Authenticate(controller.CreateProjection, 1)).Methods("POST")
	router.Handle("/api/projections", Authenticate(controller.UpdateProjection, 1)).Methods("PUT")
	router.Handle("/api/projections/{id}", Authenticate(controller.DeleteProjection, 1)).Methods("DELETE")

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
