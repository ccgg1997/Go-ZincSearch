package gateway

import (
	"github.com/ccgg1997/Go-ZincSearch/email/models"
)

type EmailGatewayIn interface {
	Store(email *models.CreateEmailCMD) (*models.CreateEmailCMD, error)
	Query(query string) ([]models.Email, error)
}
