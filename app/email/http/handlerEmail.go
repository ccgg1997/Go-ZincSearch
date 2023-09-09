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
	io.WriteString(w, "Bienvenido a Zincsearch")
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
