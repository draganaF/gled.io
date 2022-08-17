package repository

import (
	"strings"
	"time"

	apicontract "github.com/draganaF/gled.io/ProjectionService/apiContract"
	"github.com/draganaF/gled.io/ProjectionService/model"
	"gorm.io/gorm"
)

type ProjectionRepository struct {
	DB *gorm.DB
}

func (projectionRepository *ProjectionRepository) Search(params *apicontract.SearchParams) *[]model.Projection {
	projections := &[]model.Projection{}

	query := projectionRepository.DB.Table("projections").Preload("Movie").Preload("CinemaHall")

	if params.Genre != -1 {
		query.Where("projections.genre = ?", params.Genre)
	}

	if params.Name != "" {
		query.Where("LOWER(projections.name) like ?", "%"+strings.ToLower(params.Name)+"%")
	}

	if params.Director != "" {
		query.Where("LOWER(projections.director) like ?", "%"+strings.ToLower(params.Director)+"%")
	}

	if params.Actor != "" {
		query.Where("LOWER(projections.actors) like ?", "%"+strings.ToLower(params.Actor)+"%")
	}

	if params.Country != "" {
		query.Where("LOWER(projections.country) like ?", "%"+strings.ToLower(params.Country)+"%")
	}

	if params.Language != "" {
		query.Where("LOWER(projections.language) like ?", "%"+strings.ToLower(params.Language)+"%")
	}

	if params.Duration > 0 {
		query.Where("projections.duration = ?", params.Duration)
	}

	if params.Year > 0 {
		query.Where("projections.year = ?", params.Year)
	}
	if !params.DateFrom.IsZero() {
		query.Where("projections.slot >= ?", params.DateFrom)
	}

	if !params.DateTo.IsZero() {
		query.Where("projections.slot <= ?", params.DateTo)
	}

	if params.CinemaHallName != "" {
		query.Where("LOWER(projections.cinema_hall.name) like ?", "%"+params.CinemaHallName+"%")
	}

	query.Find(projections)

	return projections
}

func (projectionRepository *ProjectionRepository) ReadAll() *[]model.Projection {
	projections := &[]model.Projection{}

	result := projectionRepository.DB.Preload("Movie").Preload("CinemaHall").Where("slot >= ? AND deleted = false", time.Now()).Find(projections)

	if result.RowsAffected == 0 {
		return nil
	}

	return projections
}

func (projectionRepository *ProjectionRepository) ReadProjectionById(id uint) *model.Projection {
	projection := &model.Projection{}
	result := projectionRepository.DB.Preload("Movie").Preload("CinemaHall").First(projection, id)

	if result.RowsAffected == 0 {
		return nil
	}

	return projection
}

func (projectionRepository *ProjectionRepository) ReadProjectionsByMovieId(movieId uint) *[]model.Projection {
	projections := &[]model.Projection{}

	result := projectionRepository.DB.Preload("Movie").Preload("CinemaHall").Where("movie_id = ?", movieId).Find(projections)

	if result.RowsAffected == 0 {
		return nil
	}

	return projections
}

func (projectionRepository *ProjectionRepository) Create(projection *model.Projection) *model.Projection {
	projectionRepository.DB.Create(projection)

	return projection
}

func (projectionRepository *ProjectionRepository) Update(projection *model.Projection) *model.Projection {
	projectionRepository.DB.Save(projection)

	return projection
}

func (projectionRepository *ProjectionRepository) Delete(id uint) {
	projection := projectionRepository.ReadProjectionById(id)

	projection.Deleted = true

	projectionRepository.DB.Save(projection)
}
