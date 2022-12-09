package main

import (
	"github.com/lithammer/shortuuid/v4"
	"html/template"
	"log"
	"net/http"
)

func GenerateShortURL() string {
	shortURL := BASE_URL + shortuuid.New()
	return shortURL
}

func renderUI(w http.ResponseWriter, fileName string, data *TemplateData) {
	tmpl, err := template.ParseFiles(fileName)
	if err != nil {
		log.Fatal(err)
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		log.Fatal(err)
	}
}
