package infra

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func NewPostgre() (*sql.DB, *sql.Tx, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

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
