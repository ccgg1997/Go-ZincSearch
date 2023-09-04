package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"

	"github.com/google/uuid"
)

func main() {

	//iniciar servidor de profiling en un goroutine
	go func() {
		log.Println(http.ListenAndServe("0.0.0.0:6060", nil))
	}()

	//prueba de impresion de uuid
	fmt.Printf("si estamos en main, Hola mundo, entrada numero: %s \n", uuid.New().String())

	//servidor web principal
	http.HandleFunc("/miruta", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "mi ruta personalizada")
	})
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
	fmt.Println(uuid.New().String())

	// Forzar a vaciar el búfer de salida
	os.Stdout.Sync()
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hay peticiones en ejecucion")
	io.WriteString(w, "Hola mundo")

}
