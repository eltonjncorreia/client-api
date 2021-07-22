package main

import (
	"time"

	"github.com/google/uuid"
)

type Client struct {
	UUID          uuid.UUID `json:"uuid"`
	Nome          string    `json:"nome,omitempty"`
	Endereco      string    `json:"endereco,omitempty"`
	Cadastrado_em time.Time `json:"cadastrado_em,omitempty"`
	Atualizado_em time.Time `json:"atualizado_em,omitempty"`
}
