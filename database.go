package main

import (
	"database/sql"
)

type URLModel struct{ db *sql.DB }

func (m *URLModel) NewURL(code, longURL, shortURL string) error {
	query := `INSERT INTO urls(url_code, long_url, short_url) VALUES ($1, $2, $3)`
	_, err := m.db.Exec(query, code, longURL, shortURL)
	if err != nil {
		return err
	}
	return nil
}

func (m *URLModel) GetShortURL(longURL string) (string, error) {
	query := `SELECT short_url FROM urls WHERE long_url=$1`
	row := m.db.QueryRow(query, longURL)
	var shortURL string
	err := row.Scan(&shortURL)
	if err != nil {
		return "", err
	}
	return shortURL, nil
}

func (m *URLModel) GetLongURL(code string) (string, error) {
	query := `SELECT long_url FROM urls WHERE url_code=$1`
	row := m.db.QueryRow(query, code)
	var longURL string
	err := row.Scan(&longURL)
	if err != nil {
		return "", err
	}
	return longURL, nil
}
