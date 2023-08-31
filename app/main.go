package main

import (
	"fmt"
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
	fmt.Println(uuid.New().String())
	fmt.Println("si estamos en main Hola mundo ")
	fmt.Printf("si estamos en main Hola mundos1234589 %s", uuid.New().String())

	// Forzar a vaciar el búfer de salida
	os.Stdout.Sync()

	// Bucle infinito para mantener el programa en ejecución
	select {}
}
