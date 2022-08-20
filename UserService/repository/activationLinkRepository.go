package repository

import (
	"github.com/draganaF/gled.io/UserService/model"
	"gorm.io/gorm"
)

type ActivationLinkRepository struct {
	DB *gorm.DB
}

func (activationRepository *ActivationLinkRepository) ReadLinkByText(text string) *model.ActivationLink {
	link := &model.ActivationLink{}

	result := activationRepository.DB.Where("link = ?", text).First(link)

	if result.RowsAffected == 0 || link.Activated {
		return nil
	}

	return link
}

func (activationRepository *ActivationLinkRepository) ActivateUser(text string) *model.ActivationLink {
	link := activationRepository.ReadLinkByText(text)

	if link == nil {
		return nil
	}

	link.Activated = true

	activationRepository.DB.Save(link)

	return link
}

func (activationRepository *ActivationLinkRepository) Create(activationLink *model.ActivationLink) *model.ActivationLink {
	activationRepository.DB.Create(activationLink)

	return activationLink
}
