package service

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	apicontract "github.com/draganaF/gled.io/ProjectionService/apiContract"
	"github.com/draganaF/gled.io/ProjectionService/model"
	"github.com/draganaF/gled.io/ProjectionService/repository"
	"github.com/draganaF/gled.io/ProjectionService/utils"
)

type ProjectionService struct {
	repository *repository.ProjectionRepository
}

func NewProjectionService() *ProjectionService {
	return &ProjectionService{
		repository: &repository.ProjectionRepository{
			DB: utils.DB,
		},
	}
}

func (projectionService *ProjectionService) Search(params *apicontract.SearchParams) (*[]model.Projection, error) {
	projections := projectionService.repository.Search(params)

	if projections == nil {
		return nil, errors.New("no projections found")
	}

	return projections, nil
}

func (projectionService *ProjectionService) ReadAll() (*[]model.Projection, error) {

	projections := projectionService.repository.ReadAll()

	if projections == nil {
		return nil, errors.New("No users found")
	}

	return projections, nil
}

func (projectionService *ProjectionService) ReadProjectionById(id uint) (*model.Projection, error) {

	projection := projectionService.repository.ReadProjectionById(id)

	if projection == nil {
		return nil, errors.New("There is no projection with ID " + strconv.FormatUint(uint64(id), 10))
	}

	return projection, nil
}

func (projectionService *ProjectionService) ReadProjectionsByMovieId(id uint) (*[]model.Projection, error) {

	projection := projectionService.repository.ReadProjectionsByMovieId(id)

	if projection == nil {
		return nil, errors.New("There is no user with ID " + strconv.FormatUint(uint64(id), 10))
	}

	return projection, nil
}

func (projectionService *ProjectionService) Create(projectionDto apicontract.CreateProjectionRequest) (*model.Projection, error) {
	projection := projectionDto.ToProjection()

	fmt.Println("Service projektijce")
	movieService := NewMovieService()
	movie := movieService.repository.ReadMovieById(projection.MovieId)
	projection.Movie = *movie

	cinemaHallService := NewCinemaHallService()
	cinemaHall := cinemaHallService.repository.ReadCinemaHallById(projection.CinemaHallId)
	projection.CinemaHall = *cinemaHall

	projections := projectionService.repository.ReadAll()

	if projections != nil {
		for _, value := range *projections {
			endDate1 := projection.Slot.Add(time.Minute * time.Duration(projection.Movie.Duration))
			endDate2 := value.Slot.Add(time.Minute * time.Duration(value.Movie.Duration))

			if projection.CinemaHallId == value.CinemaHallId {
				if projectionService.AreTwoDateRangeOverlapping(projection.Slot, endDate1, value.Slot, endDate2) {
					return nil, errors.New("time slot in that hall is not available")
				}
			}
		}
	}

	projection = projectionService.repository.Create(projection)

	if projection == nil {
		return nil, errors.New("something went wrong")
	}

	return projection, nil
}

func (projectionService *ProjectionService) Update(projectionDto apicontract.UpdateProjectionRequest) (*model.Projection, error) {
	existingProjection := projectionService.repository.ReadProjectionById(projectionDto.Id)

	if existingProjection == nil {
		return nil, errors.New("No projection found with ID " + strconv.FormatUint(uint64(projectionDto.Id), 10))
	}

	existingProjection.Movie = projectionDto.Movie
	existingProjection.Slot = projectionDto.Slot
	existingProjection.CinemaHall = projectionDto.CinemaHall

	savedProjection := projectionService.repository.Update(existingProjection)

	if savedProjection == nil {
		return nil, errors.New("something went wrong")
	}

	return savedProjection, nil
}

func (projectionService *ProjectionService) Delete(id uint) {
	projectionService.repository.Delete(id)
}

func (projectionService *ProjectionService) AreTwoDateRangeOverlapping(startDate1 time.Time, endDate1 time.Time, startDate2 time.Time, endDate2 time.Time) bool {
	if AfterOrEquals(startDate1, startDate2) && BeforeOrEquals(endDate1, endDate2) {
		return true
	}

	if BeforeOrEquals(startDate1, startDate2) && AfterOrEquals(endDate1, endDate2) {
		return true
	}

	if AfterOrEquals(startDate1, startDate2) && AfterOrEquals(endDate1, endDate2) && BeforeOrEquals(startDate1, endDate2) {
		return true
	}
	if AfterOrEquals(startDate2, startDate1) && AfterOrEquals(endDate2, endDate1) && BeforeOrEquals(startDate2, endDate1) {
		return true
	}

	return false
}

func AfterOrEquals(t1 time.Time, t2 time.Time) bool {
	return t1.After(t2) || t1.Equal(t2)
}

func BeforeOrEquals(t1 time.Time, t2 time.Time) bool {
	return t1.Before(t2) || t1.Equal(t2)
}
