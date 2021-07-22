package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func clientCreateHandler(w http.ResponseWriter, r *http.Request) {
	request, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		Error(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var client Client
	if erro = json.Unmarshal(request, &client); erro != nil {
		Error(w, http.StatusBadRequest, erro)
		return
	}

	UUID, erro := CreateClient(client)
	if erro != nil {
		Error(w, http.StatusInternalServerError, erro)
		return
	}

	JSON(w, http.StatusCreated, map[string]string{"uuid": UUID})
}

func ClientListHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscar clientes"))
}

func ClientDetailHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("detalhes usuario"))
}

func ClientUpdateHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualizar usuario"))
}

func ClientDeleteHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deletar Usuário"))
}

type ErroString struct {
	Error string `json:"erro"`
}

func JSON(w http.ResponseWriter, statusCode int, dados interface{}) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(statusCode)

	if erro := json.NewEncoder(w).Encode(dados); erro != nil {
		log.Fatal(erro)
	}
}

func Error(w http.ResponseWriter, statusCode int, erro error) {
	JSON(w, statusCode, ErroString{erro.Error()})
}
