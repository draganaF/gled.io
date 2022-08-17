package apicontract

import "github.com/draganaF/gled.io/UserService/model"

type UpdateUser struct {
	Id       uint
	Name     string
	LastName string
}

func (updateUser *UpdateUser) ToUser() *model.User {
	return &model.User{
		Id:       updateUser.Id,
		Name:     updateUser.Name,
		LastName: updateUser.LastName,
	}
}
