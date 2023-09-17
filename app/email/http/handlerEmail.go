package http

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/ccgg1997/Go-ZincSearch/email/models"
	"github.com/ccgg1997/Go-ZincSearch/email/usecase"
	_ "github.com/swaggo/http-swagger"
)

type EmailHandler struct {
	emailUsecase usecase.EmailUsecase
}

func NewEmailHandler(eu usecase.EmailUsecase) *EmailHandler {
	return &EmailHandler{
		emailUsecase: eu,
	}
}

// @Summary     verify conectivity with ZincSearch
// @Description Check connectivity with ZincSearch
// @Tags        ZincSearch
// @Accept      json
// @Produce     json
// @Success     200 {string} string "La conectividad con ZincSearch esta activa, accede por medio de las peticiones HTTP de la api de email"
// @Router      /zinconection [get]
func (eh *EmailHandler) ZincSearchHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hay peticiones en ejecucion")
	io.WriteString(w, "La conectividad con ZincSearch esta activa, accede por medio de las peticiones HTTP de la api de email")
}

// @Summary      Index in zincsearch
// @Description  Store an email entry (It´s an example to add a new email value).
// @Tags         Email
// @Accept       json
// @Produce      json
// @Param        emailData   body    EmailData    true   "Email parameters"
// @Success      200 {object} SuccessResponse
// @Router       /email [post]
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

// @Summary      Search text in zincsearch
// @Description  Perform a search based on the given query. Please note that the query is a string. Search results
// @Tags         Email
// @Accept       json
// @Produce      json
// @Param        query      body    QueryParam    true   "Search parameters"
// @Success      200 {object} SearchResult
// @Router       /query [post]
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

//estructuras para la documentacion

// QueryParam represents the structure for the search query.
// @Schema
type QueryParam struct {
	Query string `json:"query"`
}

// SearchResult represents the results of a search query.
// @Schema
type SearchResult struct {
	EmailsEncontrados string `json:"Emails Encontrados"`
}

// EmailData represents the structure for the email input.
// @Schema
type EmailData struct {
	Date    string `json:"date"`
	From    string `json:"from"`
	To      string `json:"to"`
	Subject string `json:"subject"`
	XFrom   string `json:"xfrom"`
	XTo     string `json:"xto"`
	Content string `json:"content"`
}

// SuccessResponse represents a generic success response.
// @Schema
type SuccessResponse struct {
	Message string `json:"message"`
}
