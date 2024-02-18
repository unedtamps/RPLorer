package util

import (
	"bytes"
	"encoding/json"
	"text/template"
)

type EmailConfirm struct {
	Id    string
	Name  string
	Email string
}

func ParseAccountConfirmation(em EmailConfirm) string {
	var a bytes.Buffer
	tpl := template.Must(template.ParseFiles("template/email_confirm.html"))

	err := tpl.Execute(&a, em)
	if err != nil {
		Log.Fatal("Error parsing email confirmation template: ", err)
	}
	return a.String()
}

func ParseCache[T any](data string) (*T, error) {
	var result T
	err := json.Unmarshal([]byte(data), &result)
	return &result, err
}
