package model

import "gorm.io/gorm"

type CinemaHall struct {
	gorm.Model
	Id   uint `gorm:"primarykey"`
	Name string
	Rows []Row `gorm:"many2many:cinema_hall_rows;"`
}
