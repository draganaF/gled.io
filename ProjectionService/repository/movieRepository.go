package repository

import (
	"github.com/draganaF/gled.io/ProjectionService/model"

	"gorm.io/gorm"
)

type MovieRepository struct {
	DB *gorm.DB
}

func (movieRepository *MovieRepository) ReadAll() *[]model.Movie {
	movies := &[]model.Movie{}

	result := movieRepository.DB.Where(&model.Movie{Deleted: false}).Find(movies)

	if result.RowsAffected == 0 {
		return nil
	}

	return movies
}

func (movieRepository *MovieRepository) ReadMovieById(id uint) *model.Movie {
	movie := &model.Movie{}

	result := movieRepository.DB.First(movie, id)
	if result.RowsAffected == 0 || movie.Deleted {
		return nil
	}

	return movie
}

func (movieRepository *MovieRepository) Create(movie *model.Movie) *model.Movie {
	movieRepository.DB.Create(movie)

	return movie
}

func (movieRepository *MovieRepository) Update(movie *model.Movie) *model.Movie {
	movieRepository.DB.Save(movie)

	return movie
}

func (movieRepository *MovieRepository) Delete(id uint) {
	movie := movieRepository.ReadMovieById(id)

	movie.Deleted = true

	movieRepository.DB.Save(movie)
}
