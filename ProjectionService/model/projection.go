package model

import (
	"time"

	"gorm.io/gorm"
)

type Projection struct {
	gorm.Model
	Id         uint `gorm:"primarykey"`
	Movie      Movie
	Price      int
	Slot       time.Time
	CinemaHall CinemaHall
	Deleted    bool

	MovieId      uint
	CinemaHallId uint
}
