package chron

import "github.com/draganaF/gled.io/ProjectionService/service"

func DeleteUnbougthReservedTickets() {
	projectionService := service.NewProjectionService()
	projectionService.DeleteUnboughtTickets()
}
