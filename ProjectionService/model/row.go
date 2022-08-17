package model

import "gorm.io/gorm"

type Row struct {
	gorm.Model
	Id    uint `gorm:"primarykey"`
	Mark  string
	Seats int
}
