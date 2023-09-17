package gateway

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/ccgg1997/Go-ZincSearch/email/models"
)

type EmailGateway struct {
	index    string
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
	url := os.Getenv("ZINC_API_URL") + "/api/" + eg.index + "/_doc"
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

func (eg *EmailGateway) SearchQuery(query *models.CreateQueryCMD) ([]byte, error) {
	queryJSON, err := json.Marshal(query)
	if err != nil {
		fmt.Println("error parseando query")
		return nil, err
	}
	url := os.Getenv("ZINC_API_URL") + "/es/" + eg.index + "/_search"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(queryJSON))
	if err != nil {
		fmt.Println("error parseando query")
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("accept", "application/json")
	req.SetBasicAuth(eg.username, eg.password)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("error en la peticion")
		return nil, errors.New("Error al realizar la búsqueda, error en la petición" + err.Error())
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("error en la respuesta de la petición")
		return nil, errors.New("error al realizar la búsqueda, estado de la petición: " + resp.Status)
	}

	var responseBody map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&responseBody); err != nil {
		return nil, err
	}

	// Verifica si existe "hits" y "hits" contiene un array
	hits, ok := responseBody["hits"].(map[string]interface{})
	if !ok {
		return nil, errors.New("no se encontró la estructura 'hits' en la respuesta")
	}

	hitsHits, ok := hits["hits"].([]interface{})
	if !ok {
		return nil, errors.New("no se encontró la estructura 'hits.hits' en la respuesta")
	}

	// Convierte la parte de "hits.hits" de nuevo a JSON
	hitsHitsJSON, err := json.Marshal(hitsHits)
	if err != nil {
		return nil, err
	}

	fmt.Println("Búsqueda realizada con éxito")
	return hitsHitsJSON, nil
}
