package utils

import (
	"github.com/joho/godotenv"
)

func SetupEnviroment() {
	godotenv.Load(".env")
}
