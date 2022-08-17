package apicontract

import "time"

type SearchParams struct {
	Name           string
	DateFrom       time.Time
	DateTo         time.Time
	Genre          int
	Actor          string
	Director       string
	CinemaHallName string
	Duration       int
	Year           int
	Language       string
	Country        string
}
