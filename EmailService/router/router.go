package router

import (
	"net/http"
	"os"

	"github.com/draganaF/gled.io/EmailService/controller"
	"github.com/gorilla/mux"
)

func HandleRequests() {
	router := mux.NewRouter()

	router.Handle("/send", controller.SendEmail).Methods("POST")

	http.ListenAndServe(os.Getenv("SERVER_PORT"), router)
}
