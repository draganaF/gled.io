package model

import (
	"gorm.io/gorm"
)

type ActivationLink struct {
	gorm.Model
	Id     uint
	UserId uint

	Link      string
	Activated bool
}
