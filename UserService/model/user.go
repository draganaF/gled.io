package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Id       uint
	Name     string
	LastName string

	Email    string
	Password string
	Role     Role

	NumberOfBoughtTickets   int
	NumberOfSoldTickets     int
	NumberOfReservedTickets int

	NegativePoints int
	Blocked        bool
	Active         bool
	Deleted        bool
}
