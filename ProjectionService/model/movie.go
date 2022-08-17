package model

import "gorm.io/gorm"

type Movie struct {
	gorm.Model
	Id       uint `gorm:"primarykey"`
	Name     string
	Plot     string
	Genre    Genre
	Picture  string
	Duration int
	Year     int

	Country  string
	Language string

	Actors   string
	Director string

	Deleted bool
}
