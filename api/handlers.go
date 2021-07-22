package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// ClientCreateHandler retorna um cliente cadastrado no banco de dados
// Deserializa o objeto enviado para objeto do tipo Struct client
// client.Create() persiste a estrutura no banco
func ClientCreateHandler(w http.ResponseWriter, r *http.Request) {
	request, err := ioutil.ReadAll(r.Body)
	if err != nil {
		Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var client Client
	if err = json.Unmarshal(request, &client); err != nil {
		Error(w, http.StatusBadRequest, err)
		return
	}

	UUID, err := client.Create()
	if err != nil {
		Error(w, http.StatusInternalServerError, err)
		return
	}

	data := map[string]string{"uuid": UUID}

	JSON(w, http.StatusCreated, data)
}

// ClientListHandler retorna a lista de todos os clientes
func ClientListHandler(w http.ResponseWriter, r *http.Request) {
	var clientList Client

	clients, err := clientList.All()
	if err != nil {
		Error(w, http.StatusInternalServerError, err)
		return
	}

	JSON(w, http.StatusOK, clients)
}

// ClientDetailHandler retorna o detalhes de um cliente
// Recebe na URL o UUID
func ClientDetailHandler(w http.ResponseWriter, r *http.Request) {
	attrs := mux.Vars(r)

	uuid := attrs["UUID"]
	if uuid == "" {
		Error(w, http.StatusBadRequest, errors.New("UUID não definido"))
		return
	}

	var client Client
	client, err := client.Get(uuid)
	if err != nil {
		Error(w, http.StatusInternalServerError, err)
		return
	}

	if client.UUID == uuid.Nil {
		Error(w, http.StatusNotFound, errors.New("não existe client com esse uuid"))
		return
	}

	JSON(w, http.StatusOK, client)
}

// ClientUpdateHandler retorna um cliente atualizado
func ClientUpdateHandler(w http.ResponseWriter, r *http.Request) {
	attrs := mux.Vars(r)

	uuid := attrs["UUID"]
	if uuid == "" {
		Error(w, http.StatusBadRequest, errors.New("UUID não definido"))
		return
	}

	request, err := ioutil.ReadAll(r.Body)
	if err != nil {
		Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var client Client
	if err = json.Unmarshal(request, &client); err != nil {
		Error(w, http.StatusBadRequest, err)
		return
	}

	client, err = client.Update(uuid, client.Nome, client.Endereco)
	if err != nil {
		Error(w, http.StatusInternalServerError, err)
		return
	}

	JSON(w, http.StatusOK, client)
}

// ClientDeleteHandler deleta um cliente
func ClientDeleteHandler(w http.ResponseWriter, r *http.Request) {
	client := Client{}
	attrs := mux.Vars(r)

	uuid := attrs["UUID"]
	if uuid == "" {
		Error(w, http.StatusBadRequest, errors.New("UUID não definido"))
		return
	}

	client, err := client.Get(uuid)
	if err != nil {
		Error(w, http.StatusInternalServerError, err)
		return
	}

	err = client.Delete(uuid)
	if err != nil {
		Error(w, http.StatusInternalServerError, err)
		return
	}

	JSON(w, http.StatusNoContent, nil)
}

type ErroString struct {
	Error string `json:"erro"`
}

// JSON adiciona cabeçalhos na resposta e serializa o objeto
func JSON(w http.ResponseWriter, statusCode int, dados interface{}) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(statusCode)

	if statusCode != 204 {
		if erro := json.NewEncoder(w).Encode(dados); erro != nil {
			log.Fatal(erro)
		}
	}
}

// Error serializa os tipos de error
// Utiliza a função JSON para adicionar cabeçalhos na resposta
func Error(w http.ResponseWriter, statusCode int, erro error) {
	JSON(w, statusCode, ErroString{erro.Error()})
}
