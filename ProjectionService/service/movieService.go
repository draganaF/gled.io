package service

import (
	"errors"
	"strconv"

	apicontract "github.com/draganaF/gled.io/ProjectionService/apiContract"
	"github.com/draganaF/gled.io/ProjectionService/model"
	"github.com/draganaF/gled.io/ProjectionService/repository"
	"github.com/draganaF/gled.io/ProjectionService/utils"
)

type MovieService struct {
	repository *repository.MovieRepository
}

func NewMovieService() *MovieService {
	return &MovieService{
		repository: &repository.MovieRepository{
			DB: utils.DB,
		},
	}
}

func (movieService *MovieService) ReadAll() (*[]model.Movie, error) {
	movies := movieService.repository.ReadAll()

	if movies == nil {
		return nil, errors.New("No users found")
	}

	return movies, nil
}

func (movieService *MovieService) ReadMovieById(id uint) (*model.Movie, error) {
	movie := movieService.repository.ReadMovieById(id)
	if movie == nil {
		return nil, errors.New("There is no user with ID " + strconv.FormatUint(uint64(id), 10))
	}

	return movie, nil
}

func (movieService *MovieService) Create(movieDto apicontract.CreateMovieRequest) (*model.Movie, error) {
	movie := movieDto.ToMovie()

	movie = movieService.repository.Create(movie)

	if movie == nil {
		return nil, errors.New("something went wrong")
	}

	return movie, nil
}

func (movieService *MovieService) Update(movie apicontract.UpdateMovieRequest) (*model.Movie, error) {
	existingMovie := movieService.repository.ReadMovieById(movie.Id)

	if existingMovie == nil {
		return nil, errors.New("something went wrong")
	}

	existingMovie.Name = movie.Name
	existingMovie.Plot = movie.Plot
	existingMovie.Actors = movie.Actors
	existingMovie.Director = movie.Director
	existingMovie.Country = movie.Country
	existingMovie.Duration = movie.Duration
	existingMovie.Genre = model.Genre(movie.Genre)
	existingMovie.Language = movie.Language
	existingMovie.Picture = movie.Picture

	existingMovie = movieService.repository.Update(existingMovie)

	return existingMovie, nil

}

func (movieService *MovieService) Delete(id uint) {
	movieService.repository.Delete(id)
}
