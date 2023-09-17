package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"

	"github.com/ccgg1997/Go-ZincSearch/api"
	_ "github.com/ccgg1997/Go-ZincSearch/docs"
	"github.com/ccgg1997/Go-ZincSearch/email/gateway"
	customHTTP "github.com/ccgg1997/Go-ZincSearch/email/http"
	"github.com/ccgg1997/Go-ZincSearch/email/usecase"
	script "github.com/ccgg1997/Go-ZincSearch/script2"
)

//	@title			Swagger Example API
//	@version		1.0
//	@description	This is a sample server celler server.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		localhost:8080
//	@BasePath	/

//	@securityDefinitions.basic	BasicAuth

// @externalDocs.description	OpenAPI
// @externalDocs.url			https://swagger.io/resources/open-api/
func main() {

	// Iniciar servidor de profiling en un goroutine
	go func() {
		log.Println(http.ListenAndServe("0.0.0.0:6060", nil))
	}()

	go func() {
		ingesta := script.IngestaDeDatos()
		if ingesta != nil {
			fmt.Println("<100> Se realizó la ingesta de datos")
		}
		fmt.Println("<101> Ya existen los datos. No se realizó la ingesta de datos")
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
