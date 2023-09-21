package api

import (
	_ "github.com/ccgg1997/Go-ZincSearch/docs"
	customHTTP "github.com/ccgg1997/Go-ZincSearch/email/http"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	httpSwagger "github.com/swaggo/http-swagger"
	"github.com/go-chi/cors"
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

	corsConfig := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, 
	})
	r.Use(corsConfig.Handler)

	r.Post("/query", EmailHandler.QueryHandler)
	r.Get("/zinconection", EmailHandler.ZincSearchHandler)

	r.Post("/email", EmailHandler.CreateEmailHandler)

	r.Get("/swagger/*", httpSwagger.WrapHandler)
	return r
}
