package apicontract

import "github.com/draganaF/gled.io/UserService/model"

type CreateUser struct {
	Name     string
	LastName string
	Email    string
	Password string
	Role     int
}

func (createUser *CreateUser) ToUser() *model.User {
	return &model.User{
		Name:     createUser.Name,
		LastName: createUser.LastName,

		Email:    createUser.Email,
		Password: createUser.Password,
		Role:     model.Role(createUser.Role),

		NumberOfBoughtTickets:   0,
		NumberOfSoldTickets:     0,
		NumberOfReservedTickets: 0,
		NegativePoints:          0,

		Total:   0.0,
		Blocked: false,
		Deleted: false,
		Active:  false,
	}
}
