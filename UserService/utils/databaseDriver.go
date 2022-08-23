package utils

import (
	"fmt"
	"log"
	"os"

	model "github.com/draganaF/gled.io/UserService/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB
var err error
var (
	users = []model.User{
		{
			Email: "admin@gmail.com", Password: "123", Name: "Admin", LastName: "Adminovic", Role: model.Administrator, Deleted: false, Total: 0.0, Blocked: false, Active: true, NumberOfBoughtTickets: 0, NumberOfSoldTickets: 0, NumberOfReservedTickets: 0, NegativePoints: 0,
		},
		{
			Email: "radnik@email.com", Password: "123", Name: "Radnik", LastName: "Radnikovic", Role: model.Worker, Deleted: false, Total: 0.0, Blocked: false, Active: true, NumberOfBoughtTickets: 0, NumberOfSoldTickets: 1, NumberOfReservedTickets: 0, NegativePoints: 0,
		},
		{
			Email: "filip.ovic.dada@gmail.com", Password: "123", Name: "Kupac", LastName: "Kupacovic", Role: model.RegisteredUser, Deleted: false,
			Total: 0.0, Blocked: false, Active: true, NumberOfBoughtTickets: 1, NumberOfSoldTickets: 0, NumberOfReservedTickets: 0, NegativePoints: 0,
		},
		{
			Email: "filipovic.dada@gmail.com", Password: "123", Name: "Dada", LastName: "Filipovic", Role: model.RegisteredUser, Deleted: false, Total: 1000.0, Blocked: false, Active: true, NumberOfBoughtTickets: 10, NumberOfSoldTickets: 0, NumberOfReservedTickets: 3, NegativePoints: 2,
		},
	}
)

func ConnectToDatabase() {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_ADDRESS"), os.Getenv("DB_NAME"))

	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Connection to DB successfull.")
	}

	DB.Migrator().DropTable("users")
	DB.AutoMigrate(&model.User{})
	DB.AutoMigrate(&model.ActivationLink{})

	for _, user := range users {
		DB.Create(&user)
	}
}
