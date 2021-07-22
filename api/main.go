package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Rotas
	http.HandleFunc("/clientes", clientCreateHandler)

	// server
	http.ListenAndServe(":8000", nil)
}
