package repository

import (
	"github.com/draganaF/gled.io/ProjectionService/model"
	"gorm.io/gorm"
)

type CinemaHallRepository struct {
	DB *gorm.DB
}

func (cinemaHallRepository *CinemaHallRepository) ReadAll() *[]model.CinemaHall {
	halls := &[]model.CinemaHall{}

	result := cinemaHallRepository.DB.Find(halls)

	if result.RowsAffected == 0 {
		return nil
	}

	return halls
}

func (cinemaHallRepository *CinemaHallRepository) ReadCinemaHallById(id uint) *model.CinemaHall {
	cinemaHall := &model.CinemaHall{}
	result := cinemaHallRepository.DB.First(cinemaHall, id)

	if result.RowsAffected == 0 {
		return nil
	}

	return cinemaHall
}
