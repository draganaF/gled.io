package model

import "gorm.io/gorm"

type CinemaHall struct {
	gorm.Model
	Id   uint
	Name string
	Rows []Row
}
