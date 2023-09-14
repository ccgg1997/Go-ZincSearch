package models

import (
	"testing"
	"time"
)

func NewEmail(from string, subject string, to string, xfrom string, xto string, content string, date string, folder string) *CreateEmailCMD {
	return &CreateEmailCMD{
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

func Test_correctDateEmail(t *testing.T) {
	r := NewEmail(
		"camilo", "test", "camilo", "camilo", "camilo", "test", time.Now().AddDate(0, 0, -2).Format("2006-01-02"), "local/")
	err := r.Validate()

	if err != nil {
		t.Error("error:_la validacion no pasó" + err.Error())
		t.Fail()
	}
}

func Test_wrongDateEmail(t *testing.T) {
	r := NewEmail(
		"camilo", "test", "camilo", "camilo", "camilo", "test", time.Now().AddDate(0, 0, +5).Format("2006-01-02"), "local/")
	err := r.Validate()

	if err == nil {
		t.Error("error:_test with wrong date failed" + err.Error())
		t.Fail()
	}
}
