package main

import (
	"github.com/draganaF/gled.io/UserService/router"
	"github.com/draganaF/gled.io/UserService/utils"
)

func main() {
	utils.SetupEnviroment()
	utils.ConnectToDatabase()
	router.HandleRequests()
}
