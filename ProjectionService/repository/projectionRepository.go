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

	query := projectionRepository.DB.Table("projections").Preload("Movie").Preload("CinemaHall").Joins("JOIN movies m ON projections.movie_id = m.id")

	if params.Genre != -1 {
		query.Where("m.genre = ?", params.Genre)
	}

	if params.Name != "" {
		query.Where("LOWER(m.name) like ?", "%"+strings.ToLower(params.Name)+"%")
	}

	if params.Director != "" {
		query.Where("LOWER(m.director) like ?", "%"+strings.ToLower(params.Director)+"%")
	}

	if params.Actor != "" {
		query.Where("LOWER(m.actors) like ?", "%"+strings.ToLower(params.Actor)+"%")
	}

	if params.Country != "" {
		query.Where("LOWER(m.country) like ?", "%"+strings.ToLower(params.Country)+"%")
	}

	if params.Language != "" {
		query.Where("LOWER(m.language) like ?", "%"+strings.ToLower(params.Language)+"%")
	}

	if params.Duration > 0 {
		query.Where("m.duration = ?", params.Duration)
	}

	if params.Year > 0 {
		query.Where("m.year = ?", params.Year)
	}

	if !params.DateFrom.IsZero() {
		query.Where("projections.slot >= ?", params.DateFrom)
	}

	if !params.DateTo.IsZero() {
		query.Where("projections.slot <= ?", params.DateTo)
	}

	if params.CinemaHallName != "" {
		query.Joins("JOIN cinema_halls ch ON projections.cinema_hall_id = ch.id").Where("LOWER(ch.name) like ?", "%"+strings.ToLower(params.CinemaHallName)+"%")
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

func (projectionRepository *ProjectionRepository) ReadProjectionsThatStartInHalfAnHoure() *[]model.Projection {
	projections := &[]model.Projection{}
	timeNow := time.Now()
	timeInHalfAnHour := timeNow.Add(time.Minute * 30)
	result := projectionRepository.DB.Where("slot <= ? AND slot >= ?", timeInHalfAnHour, timeNow).Find(projections)

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
