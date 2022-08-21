package router

import (
	"net/http"
	"os"

	"github.com/draganaF/gled.io/ProjectionService/chron"
	"github.com/draganaF/gled.io/ProjectionService/controller"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"gopkg.in/robfig/cron.v2"
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
	router.Handle("/api/projections/search", controller.SearchProjections).Methods("POST")

	router.Handle("/api/cinema-halls", controller.GetAllCinemaHalls).Methods("GET")
	router.Handle("/api/cinema-halls/{id}", controller.GetCinemaHallById).Methods("GET")

	c := cron.New()
	c.AddFunc("@every 30m", chron.DeleteUnbougthReservedTickets)
	c.Start()

	corsOpts := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedHeaders: []string{"*"},
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
			http.MethodOptions,
			http.MethodHead,
		},
	})
	http.ListenAndServe(os.Getenv("PORT"), corsOpts.Handler(router))
}
