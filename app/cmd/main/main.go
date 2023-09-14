package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"

	"github.com/ccgg1997/Go-ZincSearch/api"
	"github.com/ccgg1997/Go-ZincSearch/email/gateway"
	customHTTP "github.com/ccgg1997/Go-ZincSearch/email/http"
	"github.com/ccgg1997/Go-ZincSearch/email/usecase"
	"github.com/ccgg1997/Go-ZincSearch/scripts"
)

func main() {

	// Iniciar servidor de profiling en un goroutine
	go func() {
		log.Println(http.ListenAndServe("0.0.0.0:6060", nil))
	}()

	go func() {
		if true {
			fmt.Println(scripts.IngestaDeDatos())
			// fmt.Println("h")
		}
	}()
	// Crear una instancia de EmailGateway
	emailGateway := gateway.NewEmailGateway("email")
	// Crear una instancia de EmailUsecase
	emailUsecase := usecase.NewEmailUsecase(*emailGateway)
	// Crear una instancia de EmailHandler
	emailHandler := customHTTP.NewEmailHandler(*emailUsecase)

	// Iniciar servidor web
	mux := api.Routes(emailHandler)
	server := api.NewServer(mux)
	server.Run()

	// Forzar a vaciar el búfer de salida
	os.Stdout.Sync()
}
