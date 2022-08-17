package router

import (
	"log"
	"net/http"
	"os"

	"github.com/draganaF/gled.io/UserService/controller"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func HandleRequests() {
	router := mux.NewRouter()

	router.Handle("/api/auth", controller.Auth).Methods("POST")
	router.Handle("/api/users/activate/{link}", controller.ActivateUser).Methods("GET")

	router.Handle("/api/users/{id}", Authenticate(controller.ReadUserById)).Methods("GET")
	router.Handle("/api/users/{email}", Authenticate(controller.ReadUserByEmail)).Methods("GET")
	router.Handle("/api/users", controller.CreateUser).Methods("POST")
	router.Handle("/api/users/update", Authenticate(controller.UpdateUser)).Methods("PUT")
	router.Handle("/api/users/block-user/{id}", Authenticate(controller.BlockUser)).Methods("GET")
	router.Handle("/api/users/delete-user/{id}", Authenticate(controller.DeleteUser)).Methods("DELETE")
	router.Handle("/api/users/search", Authenticate(controller.Search)).Methods("POST")

	router.Handle("/api/users/increment-negative-points", controller.IncrementNegativePoints).Methods("GET")
	router.Handle("/api/users/increment-bought-tickets", controller.IncrementNumberOfBougthTickets).Methods("GET")
	router.Handle("/api/users/increment-reserved-tickets", controller.IncrementNumberOfReservedTickets).Methods("GET")
	router.Handle("/api/users/increment-sold-tickets", controller.IncrementNumberOfSoldTickets).Methods("GET")

	router.Handle("/api/users/buy-tickets", controller.BuyTicket).Methods("POST")
	router.Handle("/api/users/update-ballans", controller.UpdateBallans).Methods("POST")

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
	log.Fatal(http.ListenAndServe(os.Getenv("PORT"), router), corsOpts)
}
