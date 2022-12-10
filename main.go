package main

import (
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/v5/stdlib"
	"log"
	"net/http"
)

type TemplateData struct {
	URL      string
	ShortURL string
	Error    string
}

type Application struct {
	model *URLModel
}

func main() {

	dsn := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		DB_USERNAME, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Fatal(err)
		return
	}

	app := Application{
		model: &URLModel{db: db},
	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(db)

	mux := http.NewServeMux()
	mux.HandleFunc("/", app.IndexHandler)
	mux.HandleFunc("/favicon.ico", app.FaviconHandler)
	mux.HandleFunc("/home", app.HomeHandler)
	mux.HandleFunc("/short", app.ShortHandler)

	fmt.Println("Starting server on port :4000 (http://localhost:4000)")
	log.Fatal(http.ListenAndServe(":4000", mux))

}
