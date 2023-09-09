package usecase

import (
	"github.com/ccgg1997/Go-ZincSearch/email/gateway"
	"github.com/ccgg1997/Go-ZincSearch/email/models"
)

type EmailUsecase struct {
	emailGateway gateway.EmailGateway
}

func NewEmailUsecase(eg gateway.EmailGateway) *EmailUsecase {
    return &EmailUsecase{
        emailGateway: eg,
    }
}

func (eu *EmailUsecase) CreateAndStoreEmail(cmd *models.CreateEmailCMD) (*models.CreateEmailCMD, error) {
	// Validar el email
	err := cmd.Validate()
	if err != nil {
		return nil, err
	}

	// Almacenar el email usando el gateway
	return eu.emailGateway.Store(cmd)
}
