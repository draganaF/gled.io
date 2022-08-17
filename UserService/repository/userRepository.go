package repository

import (
	"strings"

	apicontract "github.com/draganaF/gled.io/UserService/apiContract"
	model "github.com/draganaF/gled.io/UserService/model"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func (userRepository *UserRepository) ReadAll() *[]model.User {
	users := &[]model.User{}

	result := userRepository.DB.Where(&model.User{Deleted: false}).Find(&users)

	if result.RowsAffected == 0 {
		return nil
	}

	return users
}

func (userRepository *UserRepository) ReadUserById(id uint) *model.User {
	user := &model.User{}

	result := userRepository.DB.First(user, id)
	if result.RowsAffected == 0 || user.Blocked || user.Deleted {
		return nil
	}

	return user
}

func (userRepository *UserRepository) ReadUserByEmail(email string) *model.User {
	user := &model.User{}

	result := userRepository.DB.Where("email = ?", email).First(user)

	if result.RowsAffected == 0 || user.Blocked || user.Deleted {
		return nil
	}

	return user
}

func (userRepository *UserRepository) Search(searchParams *apicontract.SearchParams) *[]model.User {
	users := &[]model.User{}
	query := userRepository.DB.Table("users").Where("users.deleted = false")

	if searchParams.Role != -1 {
		query.Where("users.role = ?", searchParams.Role)
	}

	if searchParams.LastName != "" {
		query.Where("LOWER(users.last_name) like ?", "%"+strings.ToLower(searchParams.LastName)+"%")
	}

	if searchParams.Name != "" {
		query.Where("LOWER(users.name) like ?", "%"+strings.ToLower(searchParams.Name)+"%")
	}

	if searchParams.Email != "" {
		query.Where("LOWER(users.email) like ?", "%"+strings.ToLower(searchParams.Email)+"%")
	}

	query.Where("users.number_of_bought_tickets >= ? ", searchParams.FromNumberOfBoughtTickets)
	if searchParams.ToNumberOfBoughtTickets > 0 {
		query.Where("users.number_of_bought_tickets <= ? ", searchParams.ToNumberOfBoughtTickets)
	}

	query.Where("users.number_of_sold_tickets >= ?", searchParams.FromNumberOfSoldTickets)
	if searchParams.ToNumberOfSoldTickets > 0 {
		query.Where("users.number_of_sold_tickets <= ? ", searchParams.ToNumberOfSoldTickets)
	}

	query.Where("users.number_of_reserved_tickets >= ?", searchParams.FromNumberOfResevedTickets)
	if searchParams.ToNumberOfResevedTickets > 0 {
		query.Where("users.number_of_reserved_tickets <= ? ", searchParams.ToNumberOfResevedTickets)
	}

	query.Find(users)

	return users
}

func (userRepository *UserRepository) Create(user *model.User) *model.User {
	userRepository.DB.Create(user)

	return user
}

func (userRepository *UserRepository) Update(user *model.User) *model.User {
	userRepository.DB.Save(user)

	return user
}

func (userRepository *UserRepository) BlockUser(id uint) *model.User {
	user := userRepository.ReadUserById(id)

	if user == nil {
		return nil
	}

	user.Blocked = true

	userRepository.DB.Save(user)

	return user
}

func (userRepository *UserRepository) Delete(id uint) {
	user := userRepository.ReadUserById(id)

	user.Deleted = true

	userRepository.DB.Save(user)
}

func (userRepository *UserRepository) ActivateUser(id uint) *model.User {
	user := userRepository.ReadUserById(id)

	if user == nil {
		return nil
	}

	user.Active = true

	userRepository.DB.Save(user)

	return user
}
