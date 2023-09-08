package models

import (
	"testing"
	"time"
)

func NewEmail(from string, subject string, to string, xfrom string, xto string, content string, date time.Time) *CreateEmailCMD {
	return &CreateEmailCMD{
		Date:    date,
		From:    from,
		To:      to,
		Subject: subject,
		XFrom:   xfrom,
		XTo:     xto,
		Content: content,
	}
}

func Test_correctDateEmail(t *testing.T) {
	r := NewEmail(
		"camilo", "test", "camilo", "camilo", "camilo", "test", time.Now().AddDate(0, 0, -1))
	err := r.Validate()

	if err != nil {
		t.Error("error:_la validacion no pasó")
		t.Fail()
	}
}

func Test_wrongDateEmail(t *testing.T) {
	r := NewEmail(
		"camilo", "test", "camilo", "camilo", "camilo", "test", time.Now().AddDate(0, 0, +5))
	err := r.Validate()

	if err == nil {
		t.Error("error:_test with wrong date failed")
		t.Fail()
	}
}
