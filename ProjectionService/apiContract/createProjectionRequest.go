package apicontract

import (
	"time"

	"github.com/draganaF/gled.io/ProjectionService/model"
)

type CreateProjectionRequest struct {
	Movie      model.Movie
	Slot       time.Time
	CinemaHall model.CinemaHall
}

func (createProjectionRequest *CreateProjectionRequest) ToProjection() *model.Projection {
	return &model.Projection{
		Movie:      createProjectionRequest.Movie,
		Slot:       createProjectionRequest.Slot,
		CinemaHall: createProjectionRequest.CinemaHall,
		Deleted:    false,
	}
}
