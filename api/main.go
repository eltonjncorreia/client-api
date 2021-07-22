package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	// carrega arquivo .env
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Rotas
	r := mux.NewRouter()
	r.HandleFunc("/clientes", ClientListHandler).Methods("GET")
	r.HandleFunc("/clientes", ClientCreateHandler).Methods("POST")
	r.HandleFunc("/clientes/{UUID}", ClientDetailHandler).Methods("GET")
	r.HandleFunc("/clientes/{UUID}", ClientUpdateHandler).Methods("PUT")
	r.HandleFunc("/clientes/{UUID}", ClientDeleteHandler).Methods("DELETE")

	// server
	http.ListenAndServe(":8000", r)
}
