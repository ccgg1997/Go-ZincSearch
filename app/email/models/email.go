package models

import (
	"errors"
	"time"
)

type Email struct {
	Date    string
	From    string
	To      string
	Subject string
	XFrom   string
	XTo     string
	Content string
	Folder  string
}

type CreateEmailCMD struct {
	Date    string `json:"date"`
	From    string `json:"from"`
	To      string `json:"to"`
	Subject string `json:"subject"`
	XFrom   string `json:"xfrom"`
	XTo     string `json:"xto"`
	Content string `json:"content"`
	Folder  string `json:"folder"`
}

func (e *CreateEmailCMD) Validate() error {
	// Parsear el campo Date a un objeto time.Time
	layout := "2006-01-02"
	parsedDate, err := time.Parse(layout, e.Date)
	if err != nil {
		return errors.New("invalid date format" + err.Error())
	}

	// Verificar si la fecha está en el futuro
	if parsedDate.After(time.Now()) {
		return errors.New("date cannot be in the future")
	}
	return nil
}
