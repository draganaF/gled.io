package service

import (
	"errors"
	"strconv"

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
		return nil, errors.New("There is no user with ID " + strconv.FormatUint(uint64(id), 10))
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
