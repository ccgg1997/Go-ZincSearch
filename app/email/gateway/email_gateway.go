package gateway

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/ccgg1997/Go-ZincSearch/email/models"
	"net/http"
	"os"
)

type EmailGateway struct {
	index      string
	username string
	password string
}

func NewEmailGateway(index string) *EmailGateway {
	return &EmailGateway{
		index:    index,
		username: os.Getenv("ZINC_FIRST_ADMIN_USER"),
		password: os.Getenv("ZINC_FIRST_ADMIN_PASSWORD"),
	}
}

func (eg *EmailGateway) Store(email *models.CreateEmailCMD) (*models.CreateEmailCMD, error) {
	emailJSON, err := json.Marshal(email)
	if err != nil {
		return nil, err
	}
	url := os.Getenv("ZINC_API_URL")+"/api/" + eg.index + "/_doc"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(emailJSON))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("accept", "application/json")
	req.SetBasicAuth(eg.username, eg.password)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("error al almacenar el email en ZincSearch")
	}

	return email, nil
}
