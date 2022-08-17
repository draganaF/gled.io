package service

import (
	"errors"
	"strconv"

	"github.com/draganaF/gled.io/ProjectionService/model"
	"github.com/draganaF/gled.io/ProjectionService/repository"
	"github.com/draganaF/gled.io/ProjectionService/utils"
)

type CinemaHallService struct {
	repository *repository.CinemaHallRepository
}

func NewCinemaHallService() *CinemaHallService {
	return &CinemaHallService{
		repository: &repository.CinemaHallRepository{
			DB: utils.DB,
		},
	}
}

func (cinemaHallService *CinemaHallService) ReadAll() (*[]model.CinemaHall, error) {
	halls := cinemaHallService.repository.ReadAll()

	if halls == nil {
		return nil, errors.New("No cinema halls found")
	}

	return halls, nil
}

func (cinemaHallService *CinemaHallService) ReadCinemaHallById(id uint) (*model.CinemaHall, error) {
	hall := cinemaHallService.repository.ReadCinemaHallById(id)

	if hall == nil {
		return nil, errors.New("There is no cinema hall with ID" + strconv.FormatUint(uint64(id), 10))
	}

	return hall, nil
}
