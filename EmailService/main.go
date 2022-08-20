package main

import (
	"github.com/draganaF/gled.io/EmailService/router"
	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load(".env")

	router.HandleRequests()
}
