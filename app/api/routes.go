package api

import (
	_ "github.com/ccgg1997/Go-ZincSearch/docs"
	customHTTP "github.com/ccgg1997/Go-ZincSearch/email/http"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @Summary		search text in zincsearch
// @Tags			email
// @Accept			json
// @Produce		json
// @Param			param_name	path		string	true	"Descripción del parámetro"
// @Success		200			{object}	EMAILS
// @Router			/email/query [post]
func Routes(EmailHandler *customHTTP.EmailHandler) *chi.Mux {
	r := chi.NewMux()

	// Middleware
	r.Use(middleware.Logger, middleware.Recoverer)

	r.Post("/query", EmailHandler.QueryHandler)
	r.Get("/zinconection", EmailHandler.ZincSearchHandler)

	r.Post("/email", EmailHandler.CreateEmailHandler)

	r.Get("/swagger/*", httpSwagger.WrapHandler)
	return r
}
