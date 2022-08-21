package service

import (
	"errors"

	"github.com/draganaF/gled.io/UserService/model"
	"github.com/draganaF/gled.io/UserService/repository"
	"github.com/draganaF/gled.io/UserService/utils"
)

type ActivationService struct {
	repository *repository.ActivationLinkRepository
}

func NewActivationService() *ActivationService {
	return &ActivationService{
		repository: &repository.ActivationLinkRepository{
			DB: utils.DB,
		},
	}
}

func (activationService *ActivationService) GenerateActivationLink(userId uint) (*model.ActivationLink, error) {
	link := &model.ActivationLink{
		Link:      utils.GetRandomText(),
		UserId:    userId,
		Activated: false,
	}

	activationService.repository.Create(link)

	return link, nil
}

func (activationService *ActivationService) ActivateUser(text string) (*model.ActivationLink, error) {
	link := activationService.repository.ActivateUser(text)

	if link == nil {
		return nil, errors.New("Activation link is not correct")
	}

	return link, nil
}
