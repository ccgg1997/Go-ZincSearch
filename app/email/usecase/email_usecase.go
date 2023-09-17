package usecase

import (
	"encoding/json"
	"fmt"

	"github.com/ccgg1997/Go-ZincSearch/email/gateway"
	"github.com/ccgg1997/Go-ZincSearch/email/models"
)

type EmailUsecase struct {
	emailGateway gateway.EmailGateway
}

type QueryJSONData struct {
	Data []struct {
		Info struct {
			Date    string `json:"date"`
			From    string `json:"from"`
			To      string `json:"to"`
			Subject string `json:"subject"`
			XFrom   string `json:"xfrom"`
			XTo     string `json:"xto"`
			Content string `json:"content"`
			Folder  string `json:"folder"`
		} `json:"_source"`
	} `json:"data"`
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

func (eu *EmailUsecase) SentQuery(cmdquery *models.CreateQueryCMD) ([]models.CreateEmailCMD, error) {
	var EstuctEmailsFound QueryJSONData

	// Almacenar el email usando el gateway
	response, err := eu.emailGateway.SearchQuery(cmdquery)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal([]byte(response), &EstuctEmailsFound.Data); err != nil {
		fmt.Println("Error al analizar el JSON:", err)
		return nil, err
	}

	// Usamos un mapa para almacenar contenidos únicos
	contentSet := make(map[string]struct{})
	var emailsFound []models.CreateEmailCMD

	for _, item := range EstuctEmailsFound.Data {
		content := item.Info.Content

		// Verificar si el contenido ya está en el conjunto
		_, exists := contentSet[content]
		if exists {
			// El contenido ya existe, no lo añadimos nuevamente
			continue
		}

		// Añadir el contenido al conjunto y al slice
		contentSet[content] = struct{}{}
		email := models.CreateEmailCMD{
			Date:    item.Info.Date,
			From:    item.Info.From,
			To:      item.Info.To,
			Subject: item.Info.Subject,
			XFrom:   item.Info.XFrom,
			XTo:     item.Info.XTo,
			Content: item.Info.Content,
			Folder:  item.Info.Folder,
		}
		emailsFound = append(emailsFound, email)
	}

	// Imprimir el arreglo de emails únicos
	fmt.Println(emailsFound)
	return emailsFound, nil
}
