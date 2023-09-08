package models

import (
	"time"
	"errors"
)

type Email struct {
	Date    time.Time
	From    string
	To      string
	Subject string
	XFrom   string
	XTo     string
	Content string
}

type CreateEmailCMD struct {
	Date    time.Time `json:"date"`
	From    string    `json:"from"`
	To      string    `json:"to"`
	Subject string    `json:"subject"`
	XFrom   string    `json:"xfrom"`
	XTo     string    `json:"xto"`
	Content string    `json:"content"`
}

func (e *CreateEmailCMD) Validate() error {
	// check if the date that is not in the future
	if e.Date.After(time.Now()) {
		return errors.New("date cannot be in the future")
	}
	return nil
}

