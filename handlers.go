package main

import (
	"github.com/lithammer/shortuuid/v4"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func (app *Application) IndexHandler(w http.ResponseWriter, r *http.Request) {
	urlCode := r.URL.Path[1:]
	longURL, err := app.model.GetLongURL(urlCode)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	http.Redirect(w, r, longURL, http.StatusFound)
}

func (app *Application) FaviconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "ui/assets/favicon.ico")
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
	err := r.ParseForm()
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Println(err)
		return
	}

	formURL := r.PostForm.Get("url")

	if formURL == "" {
		renderUI(w, "./ui/index.gohtml",
			&TemplateData{URL: formURL, Error: "Error: url body cannot be empty"})
		return
	}

	_, err = url.ParseRequestURI(formURL)

	if err != nil {
		renderUI(w, "./ui/index.gohtml",
			&TemplateData{URL: formURL, Error: "Error: url is invalid"})
		return
	}

	uuid := shortuuid.New()
	code := uuid[:8]
	shortURL := BASE_URL + code

	err = app.model.NewURL(code, formURL, shortURL)

	if err != nil && strings.Contains(err.Error(), "duplicate key") {
		shortURL, err = app.model.GetShortURL(formURL)
		if err != nil {
			log.Println(err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
	} else if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	renderUI(w, "./ui/index.gohtml", &TemplateData{URL: formURL, ShortURL: shortURL})

}
