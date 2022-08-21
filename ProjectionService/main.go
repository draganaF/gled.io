package main

import (
	"github.com/draganaF/gled.io/ProjectionService/router"
	"github.com/draganaF/gled.io/ProjectionService/utils"
)

func main() {
	utils.SetupEnviroment()
	utils.ConnectToDatabase()
	router.HandleRequests()
}
