package utils

import (
	"fmt"
	"log"
	"os"

	"github.com/draganaF/gled.io/ProjectionService/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB
var err error

var (
	movies = []model.Movie{
		{
			Name:     "Bullet Train",
			Genre:    model.Action,
			Duration: 126,
			Director: "David Leitch",
			Year:     2022,
			Language: "English",
			Actors:   "Brad Pitt, Sandra Bullock",
			Plot:     "Five assassins aboard a fast moving bullet train find out their missions have something in common.",
			Country:  "US",
			Deleted:  false,
		},
		{
			Name:     "Minnions",
			Genre:    model.Animated,
			Duration: 87,
			Director: "Kyle Balda",
			Year:     2022,
			Language: "English/Minion",
			Actors:   "Steve Carrel",
			Plot:     "The untold story of one twelve-year-old's dream to become the world's greatest supervillain.",
			Country:  "US",
			Deleted:  false,
		},
	}
)

var (
	cinemaHalls = []model.CinemaHall{
		{
			Id:   1,
			Name: "Sala 1",
			Rows: []model.Row{
				{
					Id:    1,
					Mark:  "A",
					Seats: 10,
				},
			},
		},
		{
			Id:   2,
			Name: "Sala 2",
			Rows: []model.Row{
				{
					Id:    2,
					Mark:  "A",
					Seats: 12,
				},
				{
					Id:    3,
					Mark:  "B",
					Seats: 15,
				},
			},
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
	DB.Migrator().DropTable("rows")
	DB.AutoMigrate(&model.Row{})

	DB.Migrator().DropTable("cinema_hall_rows")

	DB.Migrator().DropTable("movies")
	DB.AutoMigrate(&model.Movie{})

	for _, movie := range movies {
		DB.Create(&movie)
	}

	DB.Migrator().DropTable("cinema_halls")
	DB.AutoMigrate(&model.CinemaHall{})

	for _, hall := range cinemaHalls {
		DB.Create(&hall)
	}

	DB.Migrator().DropTable("projections")
	DB.AutoMigrate(&model.Projection{})

}
