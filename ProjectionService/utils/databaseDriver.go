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
			Actors:   []string{"Brad Pitt", "Joey King", "Logan Lerman", "Brian Tyree Henry", "Aaron Taylor-Johnson"},
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
			Actors:   []string{"Steve Carell", "Pierre Coffin", "Alan Arkin", "Taraji P. Henson", "Michelle Yeoh"},
			Plot:     "The untold story of one twelve-year-old's dream to become the world's greatest supervillain.",
			Country:  "US",
			Deleted:  false,
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

	DB.Migrator().DropTable("movies")
	DB.AutoMigrate(&model.Movie{})

	for _, movie := range movies {
		DB.Create(&movie)
	}
}
