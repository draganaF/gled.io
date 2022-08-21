package utils

import (
	"github.com/draganaF/gled.io/ProjectionService/chron"
	"gopkg.in/robfig/cron.v2"
)

func SetupChron() {
	c := cron.New()
	c.AddFunc("@every ", chron.DeleteUnbougthReservedTickets())
}
