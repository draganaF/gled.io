package model

import "gorm.io/gorm"

type Row struct {
	gorm.Model
	Mark  string
	Seats int
}
