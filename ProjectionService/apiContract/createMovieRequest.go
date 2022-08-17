package apicontract

import "github.com/draganaF/gled.io/ProjectionService/model"

type CreateMovieRequest struct {
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

func (createMovieRequest *CreateMovieRequest) ToMovie() *model.Movie {
	return &model.Movie{
		Name:     createMovieRequest.Name,
		Plot:     createMovieRequest.Plot,
		Genre:    model.Genre(createMovieRequest.Genre),
		Picture:  createMovieRequest.Picture,
		Duration: createMovieRequest.Duration,
		Year:     createMovieRequest.Year,

		Country:  createMovieRequest.Country,
		Language: createMovieRequest.Language,

		Actors:   createMovieRequest.Actors,
		Director: createMovieRequest.Director,

		Deleted: false,
	}
}
