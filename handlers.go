package main

import (
	"github.com/lithammer/shortuuid/v4"
	"html/template"
	"log"
	"net/http"
	"net/url"
)

func (app *Application) IndexHandler(w http.ResponseWriter, r *http.Request) {

}

func (app *Application) HomeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./ui/index.gohtml")
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Println(err)
		return
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Println(err)
		return
	}
}

func (app *Application) ShortHandler(w http.ResponseWriter, r *http.Request) {
	data := TemplateData{}
	err := r.ParseForm()
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Println(err)
		return
	}
	formURL := r.PostForm.Get("url")

	if formURL == "" {
		data.Error = "Error: url body cannot be empty"
		renderUI(w, "./ui/index.gohtml", &data)
		return
	}

	_, err = url.ParseRequestURI(formURL)
	if err != nil {
		data.Error = "Error: url is invalid"
		renderUI(w, "./ui/index.gohtml", &data)
		return
	}

	code := shortuuid.New()
	shortURL := BASE_URL + "/" + code
	err = app.model.NewURL(code, formURL, shortURL)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Println(err)
		return
	}

	data.ShortURL = shortURL
	data.URL = formURL
	renderUI(w, "./ui/index.gohtml", &data)

}
