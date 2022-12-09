package main

import "database/sql"

type URLModel struct{ db *sql.DB }

func (m *URLModel) NewURL(code, longURL, shortURL string) error {
	query := `INSERT INTO links(url_code, long_url, short_url) VALUES ($1, $2, $3)`
	_, err := m.db.Exec(query, code, longURL, shortURL)
	if err != nil {
		return err
	}
	return nil
}

func (m *URLModel) GetShortURL(urlCode string) (string, error) {
	query := `SELECT long_url FROM links WHERE url_code=$1`
	rows, err := m.db.Query(query, urlCode)
	if err != nil {
		return "", err
	}
	var longURL string
	err = rows.Scan(&longURL)
	if err != nil {
		return "", err
	}
	return longURL, nil
}
