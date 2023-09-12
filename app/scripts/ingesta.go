package scripts

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func IngestaDeDatos() int {
	// Consulta a ZincSearch
	exists, err := CheckIndexExists()
	if err != nil {
		log.Fatalf("Error en la consulta del index: %v", err)
	}

	if !exists {
		return 0
	}

	// Leer datos de la carpeta
	emails := readEmailData()
	if err != nil {
		log.Fatalf("Error al leer datos de email: %v", err)
	}
	fmt.Println(emails)
	CheckIndexExists()

	return emails

	// //opcion1
	// if !exists {
	// 	// Leer datos de la carpeta
	// 	emails, err := readEmailData()
	// 	if err != nil {
	// 		log.Fatalf("Error al leer datos de email: %v", err)
	// 	}

	// 	// Ingesta de datos
	// 	for _, email := range emails {
	// 		err := storeEmail(email)
	// 		if err != nil {
	// 			log.Printf("Error al almacenar email: %v", err)
	// 		}
	// 	}
	// }
}

func CheckIndexExists() (bool, error) {
	//crear peticion
	url := os.Getenv("ZINC_API_URL") + "/api/index/history"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return false, err
	}

	//datos de la peticion
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("accept", "application/json")
	req.SetBasicAuth(os.Getenv("ZINC_FIRST_ADMIN_USER"), os.Getenv("ZINC_FIRST_ADMIN_PASSWORD"))

	//enviar peticion
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return false, err
	}

	//evaluar respuesta
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return false, errors.New("error: no existe el index")
	}
	return true, nil
}

func readEmailData() int {
	root := "../../data/enron_mail_20110402/maildir/allen-p/_sent_mail" // Cambia esto a la ruta de tu carpeta principal
	var count int
	go func() {}()
	err := filepath.Walk(root,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			// Verifica si el elemento no es un directorio
			if !info.IsDir() {
				// Incrementa el contador para cada archivo (independientemente de la extensión)
				count++
			}
			return nil
		})

	if err != nil {
		fmt.Printf("Error al procesar archivos: %s\n", err)
		return 0
	}

	fmt.Printf("Se encontraron %d archivos .txt\n", count)

	return count
}

// func storeEmail(email models.CreateEmailCMD) error {
// 	data, err := json.Marshal(email)
// 	if err != nil {
// 		return err
// 	}

// 	resp, err := http.Post("www.hola.com.co", "application/json", bytes.NewBuffer(data))
// 	if err != nil {
// 		return err
// 	}
// 	defer resp.Body.Close()

// 	if resp.StatusCode != http.StatusCreated {
// 		body, _ := io.ReadAll(resp.Body)
// 		return errors.New(string(body))
// 	}

// 	return nil
// }
