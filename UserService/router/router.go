package router

import (
	"net/http"
	"os"

	"github.com/draganaF/gled.io/UserService/controller"
	"github.com/draganaF/gled.io/UserService/model"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func HandleRequests() {
	router := mux.NewRouter()

	router.Handle("/api/auth", controller.Auth).Methods("POST")
	router.Handle("/api/authorize", controller.AuthorizeUser).Methods("GET")

	router.Handle("/api/users/activate/{link}", controller.ActivateUser).Methods("GET")

	router.Handle("/api/users/buy-tickets", Authenticate(Authorize(controller.BuyTicket, model.RegisteredUser, model.Worker))).Methods("POST")
	router.Handle("/api/users/update-balance", Authenticate(Authorize(controller.UpdateBallans, model.Worker))).Methods("POST")

	router.Handle("/api/users/{id}", Authenticate(controller.ReadUserById)).Methods("GET")
	router.Handle("/api/users/by-email/{email}", Authenticate(Authorize(controller.ReadUserByEmail, model.Administrator, model.RegisteredUser, model.Worker))).Methods("GET")
	router.Handle("/api/users", controller.CreateUser).Methods("POST")
	router.Handle("/api/users/update", Authenticate(controller.UpdateUser)).Methods("PUT")
	router.Handle("/api/users/update/password", Authenticate(controller.ChangePassword)).Methods("PUT")
	router.Handle("/api/users/block-user/{id}", Authenticate(Authorize(controller.BlockUser, model.Administrator))).Methods("GET")
	router.Handle("/api/users/delete-user/{id}", Authenticate(Authorize(controller.DeleteUser, model.Administrator))).Methods("DELETE")
	router.Handle("/api/users/search", Authenticate(Authorize(controller.Search, model.Administrator, model.Worker))).Methods("POST")

	router.Handle("/api/users/increment-negative-points/{id}", controller.IncrementNegativePoints).Methods("GET")
	router.Handle("/api/users/increment-bought-tickets/{id}", Authenticate(Authorize(controller.IncrementNumberOfBougthTickets, model.RegisteredUser, model.Worker))).Methods("GET")
	router.Handle("/api/users/increment-reserved-tickets/{id}", Authenticate(Authorize(controller.IncrementNumberOfReservedTickets, model.RegisteredUser))).Methods("GET")
	router.Handle("/api/users/increment-sold-tickets/{id}", Authenticate(Authorize(controller.IncrementNumberOfSoldTickets, model.Worker))).Methods("GET")

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
