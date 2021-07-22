package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func connection() (*sql.DB, error) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Error loading .env file")
	}

	connectionString := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}

func CreateClient(c Client) (string, error) {
	var lastID string
	var stmt string

	db, err := connection()
	if err != nil {
		return "", err
	}
	defer db.Close()

	stmt = "insert into clients (nome, endereco, cadastrado_em, atualizado_em) values($1, $2, $3, $4) RETURNING uuid"

	erro := db.QueryRow(stmt, c.Nome, c.Endereco, c.Cadastrado_em, c.Atualizado_em).Scan(&lastID)
	if erro != nil {
		return "", erro
	}
	return lastID, nil
}
