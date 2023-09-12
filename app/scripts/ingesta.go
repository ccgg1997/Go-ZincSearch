package scripts

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/mail"
	"os"
	"path/filepath"

	"github.com/ccgg1997/Go-ZincSearch/email/models"
)

func IngestaDeDatos() int {
	// Consulta a ZincSearch
	exists, err := CheckIndexExists()
	if err != nil {
		log.Printf("Error en la consulta del index: %v", err)
	}

	if !exists {
		log.Printf("Se va a crear por default el indice: %v", "email")
	}

	// Leer datos de la carpeta
	emails := readEmailData()
	if err != nil {
		log.Fatalf("Error al leer datos de email: %v", err)
	}
	fmt.Println(emails)

	return emails

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
	root := "../../data/enron_mail_20110402/maildir/allen-p" // Cambia esto a la ruta de tu carpeta principal
	var count int
	err := filepath.Walk(root,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			// Verifica si el elemento no es un directorio
			if !info.IsDir() {

				// Parsea el correo electrónico
				file, err := os.Open(path)
				if err != nil {
					return err
				}
				reader := bufio.NewReader(file)
				email, err := mail.ReadMessage(reader)
				if err != nil {
					return err
				}

				from := email.Header.Get("From")
				toAddresses, _ := email.Header.AddressList("To")
				subject := email.Header.Get("Subject")
				date := email.Header.Get("Date")
				xFrom := email.Header.Get("X-From")
				xTo := email.Header.Get("X-To")
				to := ""
				if len(toAddresses) > 0 {
					to = toAddresses[0].Address
				}
				bodyByte, err := io.ReadAll(email.Body)
				if err != nil {
					return err
				}
				body := string(bodyByte)

				//crear structura
				r := NewEmail(
					date, from, to, subject, xFrom, xTo, body)
				storeEmail(*r)
				// Imprime información sobre el correo electrónico
				fmt.Println(r.Date)
				fmt.Println(r.From)
				fmt.Println(r.To)
				fmt.Println(r.XTo)
				fmt.Println(r.XFrom)
				fmt.Println(r.Content)

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

func storeEmail(email models.CreateEmailCMD) error {
	emailJSON, err := json.Marshal(email)
	if err != nil {
		return err
	}

	url := os.Getenv("ZINC_API_URL") + "/api/" + "mail" + "/_doc"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(emailJSON))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("accept", "application/json")
	req.SetBasicAuth(os.Getenv("ZINC_FIRST_ADMIN_USER"), os.Getenv("ZINC_FIRST_ADMIN_PASSWORD"))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New("error al almacenar el email en ZincSearch")
	}

	return nil
}

func NewEmail(date string, from string, to string, subject string, xfrom string, xto string, content string) *models.CreateEmailCMD {
	return &models.CreateEmailCMD{
		Date:    date,
		From:    from,
		To:      to,
		Subject: subject,
		XFrom:   xfrom,
		XTo:     xto,
		Content: content,
	}
}
