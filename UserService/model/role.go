package model

type Role int

const (
	RegisteredUser Role = iota
	Worker
	Administrator
)
