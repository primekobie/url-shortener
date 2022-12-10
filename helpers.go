package main

import (
	"html/template"
	"log"
	"net/http"
)

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
