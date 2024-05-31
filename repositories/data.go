package repositories

import (
	"context"

	"github.com/jackc/pgx/v4"
)

var db *pgx.Conn

func ConnectionPSQL(conn *pgx.Conn){
	db=conn
}

func SaveURLMapping(shortURL, originalURL string) error {
    _, err := db.Exec(context.Background(),
        "INSERT INTO urls (short_url, original_url) VALUES ($1, $2)", shortURL, originalURL)
    return err
}

func GetOriginalURL(shortURL string) (string, error) {
    var originalURL string
    err := db.QueryRow(context.Background(),
        "SELECT original_url FROM urls WHERE short_url=$1", shortURL).Scan(&originalURL)
    return originalURL, err
}