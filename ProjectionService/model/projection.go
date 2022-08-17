package model

import (
	"time"

	"gorm.io/gorm"
)

type Projection struct {
	gorm.Model
	Movie      Movie
	Slot       time.Time
	CinemaHall CinemaHall
	Deleted    bool

	MovieId      uint
	CinemaHallId uint
}
