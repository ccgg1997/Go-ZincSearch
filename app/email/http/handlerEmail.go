package http

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/ccgg1997/Go-ZincSearch/email/models"
	"github.com/ccgg1997/Go-ZincSearch/email/usecase"
)

type EmailHandler struct {
	emailUsecase usecase.EmailUsecase
}

func NewEmailHandler(eu usecase.EmailUsecase) *EmailHandler {
	return &EmailHandler{
		emailUsecase: eu,
	}
}

func (eh *EmailHandler) ZincSearchHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hay peticiones en ejecucion")
	io.WriteString(w, "La conectividad con ZincSearch esta activa, accede por medio de las peticiones HTTP de la api de email")
}

func (eh *EmailHandler) CreateEmailHandler(w http.ResponseWriter, r *http.Request) {
	var cmd models.CreateEmailCMD
	if err := json.NewDecoder(r.Body).Decode(&cmd); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	email, err := eh.emailUsecase.CreateAndStoreEmail(&cmd)
	if err != nil {
		http.Error(w, "Error, formato invalido del body", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{"creado con éxito": email})
}

func (eh *EmailHandler) QueryHandler(w http.ResponseWriter, r *http.Request) {

	var requestBody map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		return
	}

	// Accede al campo query del objeto Go
	text, ok := requestBody["query"].(string)
	if !ok {
		return
	}

	// Instanciar query
	var query models.CreateQueryCMD
	query.Query.Bool.Should = []struct {
		MatchPhrase map[string]struct {
			Query string  `json:"query"`
			Boost float64 `json:"boost"`
		} `json:"match_phrase"`
	}{
		{
			MatchPhrase: map[string]struct {
				Query string  `json:"query"`
				Boost float64 `json:"boost"`
			}{
				"content": {
					Query: text,
					Boost: 2,
				},
			},
		},
		{
			MatchPhrase: map[string]struct {
				Query string  `json:"query"`
				Boost float64 `json:"boost"`
			}{
				"date": {
					Query: text,
					Boost: 1.5,
				},
			},
		},
		{
			MatchPhrase: map[string]struct {
				Query string  `json:"query"`
				Boost float64 `json:"boost"`
			}{
				"xfrom": {
					Query: text,
					Boost: 1.6,
				},
			},
		},
		{
			MatchPhrase: map[string]struct {
				Query string  `json:"query"`
				Boost float64 `json:"boost"`
			}{
				"xto": {
					Query: text,
					Boost: 1.6,
				},
			},
		},
	}

	query.Size = 20

	email, err := eh.emailUsecase.SentQuery(&query)
	if err != nil {
		http.Error(w, "Error, formato invalido del body", http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{"Emails Encontrados": email})
}
