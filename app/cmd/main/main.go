package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"

	"github.com/ccgg1997/Go-ZincSearch/email/gateway"
	customHTTP "github.com/ccgg1997/Go-ZincSearch/email/http"
	"github.com/ccgg1997/Go-ZincSearch/email/usecase"
)

func main() {

	// Iniciar servidor de profiling en un goroutine
	go func() {
		log.Println(http.ListenAndServe("0.0.0.0:6060", nil))
	}()
	

	// Crear una instancia de EmailGateway
	emailGateway := gateway.NewEmailGateway("history")

	// Crear una instancia de EmailUsecase
	emailUsecase := usecase.NewEmailUsecase(*emailGateway)

	// Crear una instancia de EmailHandler
	emailHandler := customHTTP.NewEmailHandler(*emailUsecase)

	// Iniciar servidor web
	mux := Routes(emailHandler)
	server := NewServer(mux)
	server.Run()

	// Forzar a vaciar el búfer de salida
	os.Stdout.Sync()
}
