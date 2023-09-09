package main

import (

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	customHTTP "github.com/ccgg1997/Go-ZincSearch/email/http"
)

func Routes(emailHandler *customHTTP.EmailHandler) *chi.Mux {
	r := chi.NewMux()

	// Middleware
	r.Use(middleware.Logger, middleware.Recoverer)

	// Rutas para emails
	r.Route("/email", func(r chi.Router) {
		r.Get("/", emailHandler.ZincSearchHandler)
		r.Post("/", emailHandler.CreateEmailHandler)
	})

	return r
}

