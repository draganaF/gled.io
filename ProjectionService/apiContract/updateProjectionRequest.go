package apicontract

import (
	"time"

	"github.com/draganaF/gled.io/ProjectionService/model"
)

type UpdateProjectionRequest struct {
	Id         uint
	Movie      model.Movie
	Slot       time.Time
	CinemaHall model.CinemaHall
}
