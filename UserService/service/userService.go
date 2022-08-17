package service

import (
	"errors"
	"strconv"

	apicontract "github.com/draganaF/gled.io/UserService/apiContract"
	"github.com/draganaF/gled.io/UserService/model"
	"github.com/draganaF/gled.io/UserService/repository"
	"github.com/draganaF/gled.io/UserService/utils"
)

type UserService struct {
	repository *repository.UserRepository
}

func NewUserService() *UserService {
	return &UserService{
		repository: &repository.UserRepository{
			DB: utils.DB,
		},
	}
}

func (userService *UserService) Search(params *apicontract.SearchParams) (*[]model.User, error) {
	users := userService.repository.Search(params)

	if users == nil {
		return nil, errors.New("No users found")
	}

	return users, nil
}

func (userService *UserService) ReadAll() (*[]model.User, error) {
	users := userService.repository.ReadAll()

	if users == nil {
		return nil, errors.New("No users found")
	}

	return users, nil
}

func (userService *UserService) ReadUserById(id uint) (*model.User, error) {
	user := userService.repository.ReadUserById(id)
	if user == nil {
		return nil, errors.New("There is no user with ID " + strconv.FormatUint(uint64(id), 10))
	}

	return user, nil
}

func (userService *UserService) ReadUserByEmail(email string) (*model.User, error) {
	user := userService.repository.ReadUserByEmail(email)

	if user == nil {
		return nil, errors.New("There is no user with email " + email)
	}

	return user, nil
}

func (userService *UserService) Create(userDto apicontract.CreateUser) (*model.User, error) {

	if userService.repository.ReadUserByEmail(userDto.Email) != nil {
		return nil, errors.New("Given email " + userDto.Email + " is already in use.")
	}

	user := userDto.ToUser()

	user = userService.repository.Create(user)

	if user == nil {
		return nil, errors.New("something went wrong")
	}

	return user, nil
}

func (userService *UserService) Update(updateUser apicontract.UpdateUser) (*model.User, error) {
	existingUser := userService.repository.ReadUserById(updateUser.Id)

	if existingUser == nil {
		return nil, errors.New("No user found with ID " + strconv.FormatUint(uint64(updateUser.Id), 10))
	}

	existingUser.Name = updateUser.Name
	existingUser.LastName = updateUser.LastName

	savedUser := userService.repository.Update(existingUser)

	if savedUser == nil {
		return nil, errors.New("something went wrong")
	}

	return savedUser, nil
}

func (userService *UserService) ChangeUserPassword(request apicontract.ChangePassword) (*model.User, error) {
	user := userService.repository.ReadUserById(request.UserId)

	if user == nil {
		return nil, errors.New("User not found")
	}

	if user.Password != request.OldPassword {
		return nil, errors.New("passwords do not match")
	}

	user.Password = request.NewPassword

	userService.repository.Update(user)

	return user, nil
}

func (userService *UserService) IncrementNegativePoints(id uint) (*model.User, error) {
	existingUser := userService.repository.ReadUserById(id)

	if existingUser == nil {
		return nil, errors.New("User does not exist in system")
	}

	if existingUser.Role != model.RegisteredUser {
		return nil, errors.New("Only registered user has negative points")
	}

	existingUser.NegativePoints += 1

	if existingUser.NegativePoints == 3 {
		existingUser.Blocked = true
	}

	user := userService.repository.Update(existingUser)

	if user == nil {
		return nil, errors.New("Something went wrong")
	}

	return user, nil

}

func (userService *UserService) IncrementNumberOfReservedTickets(id uint) (*model.User, error) {
	existingUser := userService.repository.ReadUserById(id)

	if existingUser == nil {
		return nil, errors.New("No user found with ID " + strconv.FormatUint(uint64(id), 10))
	}

	existingUser.NumberOfReservedTickets += 1

	savedUser := userService.repository.Update(existingUser)

	if savedUser == nil {
		return nil, errors.New("something went wrong")
	}

	return savedUser, nil
}

func (userService *UserService) IncrementNumberOfBoughtTickets(id uint) (*model.User, error) {
	existingUser := userService.repository.ReadUserById(id)

	if existingUser == nil {
		return nil, errors.New("No user found with ID " + strconv.FormatUint(uint64(id), 10))
	}

	existingUser.NumberOfBoughtTickets += 1

	savedUser := userService.repository.Update(existingUser)

	if savedUser == nil {
		return nil, errors.New("something went wrong")
	}

	return savedUser, nil
}

func (userService *UserService) IncrementNumberOfSoldTickets(id uint) (*model.User, error) {
	existingUser := userService.repository.ReadUserById(id)

	if existingUser == nil {
		return nil, errors.New("No user found with ID " + strconv.FormatUint(uint64(id), 10))
	}

	existingUser.NumberOfSoldTickets += 1

	savedUser := userService.repository.Update(existingUser)

	if savedUser == nil {
		return nil, errors.New("something went wrong")
	}

	return savedUser, nil
}

func (userService *UserService) AddMoney(userId uint, amount float32) (*model.User, error) {
	user := userService.repository.ReadUserById(userId)

	if user == nil {
		return nil, errors.New("No user found with ID " + strconv.FormatUint(uint64(userId), 10))
	}

	user.Total += amount

	savedUser := userService.repository.Update(user)
	if savedUser == nil {
		return nil, errors.New("something went wrong")
	}

	return savedUser, nil
}

func (userService *UserService) BlockUser(id uint) (*model.User, error) {
	user := userService.repository.BlockUser(id)

	if user == nil {
		return nil, errors.New("something went wrong")
	}

	return user, nil
}

func (userService *UserService) Delete(id uint) {
	userService.repository.Delete(id)
}

func (userService *UserService) ActivateUser(id uint) (*model.User, error) {
	user := userService.repository.ActivateUser(id)

	if user == nil {
		return nil, errors.New("something went worng")
	}

	return user, nil
}
