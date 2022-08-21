package apicontract

import (
	"time"

	"github.com/draganaF/gled.io/ProjectionService/model"
)

type CreateProjectionRequest struct {
	MovieId uint
	Slot    time.Time
	HallId  uint
	Price   int
}

func (createProjectionRequest *CreateProjectionRequest) ToProjection() *model.Projection {

	return &model.Projection{
		Slot:         createProjectionRequest.Slot,
		MovieId:      createProjectionRequest.MovieId,
		CinemaHallId: createProjectionRequest.HallId,
		Price:        createProjectionRequest.Price,
		Deleted:      false,
	}
}
