package apicontract

import "github.com/draganaF/gled.io/ProjectionService/model"

type UpdateMovieRequest struct {
	Id       uint
	Name     string
	Plot     string
	Genre    int
	Picture  string
	Duration int
	Year     int

	Country  string
	Language string

	Actors   []string
	Director string
}

func (updateMovieRequest *UpdateMovieRequest) ToMovie() *model.Movie {
	return &model.Movie{
		Name:     updateMovieRequest.Name,
		Plot:     updateMovieRequest.Plot,
		Genre:    model.Genre(updateMovieRequest.Genre),
		Picture:  updateMovieRequest.Picture,
		Duration: updateMovieRequest.Duration,
		Year:     updateMovieRequest.Year,

		Country:  updateMovieRequest.Country,
		Language: updateMovieRequest.Language,

		Actors:   updateMovieRequest.Actors,
		Director: updateMovieRequest.Director,
	}
}
