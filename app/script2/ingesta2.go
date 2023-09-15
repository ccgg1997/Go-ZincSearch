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
	"sync"

	"github.com/ccgg1997/Go-ZincSearch/email/models"
)

func IngestaDeDatos() error {
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

	return nil

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

func readEmailData() []models.CreateEmailCMD {
	root := "../../data/enron_mail_20110402/maildir"
	var emails []models.CreateEmailCMD
	var errorMails []string
	ch := make(chan models.CreateEmailCMD, 14)
	var wg sync.WaitGroup

	// Walk files
	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			errorMails = append(errorMails, fmt.Sprintf("Error al acceder al archivo: %s", err))
			return nil
		}
		if !info.IsDir() {
			wg.Add(1)
			go processFile(path, ch, &wg, &errorMails)
		}
		return nil
	})

	// Close channel after all goroutines finish
	go func() {
		wg.Wait()
		close(ch)
	}()

	// Collect emails from channel
	for email := range ch {
		emails = append(emails, email)
		if len(emails) == 9000 {
			storeEmail(emails)
			emails = []models.CreateEmailCMD{}
		}
	}

	// Store any remaining emails
	if len(emails) > 0 {
		storeEmail(emails)
	}

	// Aquí puedes manejar o imprimir los correos electrónicos erróneos si lo deseas
	for _, errorMsg := range errorMails {
		fmt.Println(errorMsg)
	}

	return emails
}

func processFile(path string, ch chan models.CreateEmailCMD, wg *sync.WaitGroup, errorMails *[]string) {
	defer wg.Done()

	file, err := os.Open(path)
	if err != nil {
		*errorMails = append(*errorMails, fmt.Sprintf("Error al abrir el archivo: %s", err))
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	email, err := mail.ReadMessage(reader)
	if err != nil {
		*errorMails = append(*errorMails, fmt.Sprintf("Error al leer el correo electrónico: %s, path: %s", err, path))
		return
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
		*errorMails = append(*errorMails, fmt.Sprintf("Error al leer el cuerpo del correo electrónico: %s, path: %s", err, path))
		return
	}
	body := string(bodyByte)
	folder := email.Header.Get("X-Folder")
	fmt.Println(folder + " " + from)
	ch <- *NewEmail(date, from, to, subject, xFrom, xTo, body, folder) // Asegúrate de que la función NewEmail ahora devuelva un puntero a models.CreateEmailCMD

}

func storeEmail(emails []models.CreateEmailCMD) error {
	payload := struct {
		Index   string                  `json:"index"`
		Records []models.CreateEmailCMD `json:"records"`
	}{
		Index:   "email",
		Records: emails,
	}

	emailJSON, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	if err != nil {
		return err
	}

	url := os.Getenv("ZINC_API_URL") + "/api/" + "/_bulkv2"
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

func NewEmail(date string, from string, to string, subject string, xfrom string, xto string, content string, folder string) *models.CreateEmailCMD {
	return &models.CreateEmailCMD{
		Date:    date,
		From:    from,
		To:      to,
		Subject: subject,
		XFrom:   xfrom,
		XTo:     xto,
		Content: content,
		Folder:  folder,
	}
}
