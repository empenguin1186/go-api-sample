package infra

import (
	"database/sql"
	"fmt"
	"os"
)

func NewPostgre() (*sql.DB, *sql.Tx, error) {
	host := os.Getenv("POSTGRES_HOST")
	user := os.Getenv("POSTGRES_USER")
	pswd := os.Getenv("POSTGRES_PASSWORD")
	dbnm := os.Getenv("POSTGRES_DB")
	port := os.Getenv("POSTGRES_PORT")

	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, pswd, dbnm))
	if err != nil {
		return nil, nil, err
	}

	return db, nil, err
}
