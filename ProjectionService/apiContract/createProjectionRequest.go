package apicontract

import (
	"time"

	"github.com/draganaF/gled.io/ProjectionService/model"
)

type CreateProjectionRequest struct {
	MovieId      uint
	Slot         time.Time
	CinemaHallId uint
}

func (createProjectionRequest *CreateProjectionRequest) ToProjection() *model.Projection {
	return &model.Projection{
		Slot:         createProjectionRequest.Slot,
		MovieId:      createProjectionRequest.MovieId,
		CinemaHallId: createProjectionRequest.CinemaHallId,
		Deleted:      false,
	}
}