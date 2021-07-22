package main

import (
	"time"

	"github.com/google/uuid"
)

type Client struct {
	UUID          uuid.UUID `json:"uuid"`
	Nome          string    `json:"nome"`
	Endereco      string    `json:"endereco"`
	Cadastrado_em time.Time `json:"cadastrado_em"`
	Atualizado_em time.Time `json:"atualizado_em"`
}

func (c Client) Create() (string, error) {
	var lastID string
	var stmt string

	db, err := Connection()
	if err != nil {
		return "", err
	}
	defer db.Close()

	stmt = "insert into clients (nome, endereco, cadastrado_em, atualizado_em) values($1, $2, $3, $4) RETURNING uuid"

	err = db.QueryRow(stmt, c.Nome, c.Endereco, c.Cadastrado_em, c.Atualizado_em).Scan(&lastID)
	if err != nil {
		return "", err
	}
	return lastID, nil
}

func (c Client) Get(uuid string) (Client, error) {
	db, err := Connection()
	if err != nil {
		return c, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM clients WHERE uuid=$1", uuid)
	if err != nil {
		return c, err
	}
	defer rows.Close()

	if rows.Next() {
		erro := rows.Scan(&c.UUID, &c.Nome, &c.Endereco, &c.Cadastrado_em, &c.Atualizado_em)
		if erro != nil {
			return c, erro
		}
	}
	return c, nil
}

func (c Client) All() ([]Client, error) {
	var clientList []Client

	db, err := Connection()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM clients ORDER BY cadastrado_em DESC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var c Client

		err := rows.Scan(&c.UUID, &c.Nome, &c.Endereco, &c.Cadastrado_em, &c.Atualizado_em)
		if err != nil {
			return nil, err
		}

		clientList = append(clientList, c)
	}
	return clientList, nil
}

func (c Client) Update(uuid, Nome, Endereco string) (Client, error) {
	var lastUUID string
	stmt := "UPDATE clients SET  nome = $2, endereco = $3 WHERE uuid = $1 RETURNING uuid;"

	db, err := Connection()
	if err != nil {
		return c, err
	}
	defer db.Close()

	err = db.QueryRow(stmt, uuid, c.Nome, c.Endereco).Scan(&lastUUID)
	if err != nil {
		return c, err
	}

	client, err := c.Get(lastUUID)
	if err != nil {
		return c, err
	}

	return client, nil
}

func (c Client) Delete(uuid string) error {
	stmt := "DELETE FROM clients WHERE uuid = $1;"

	db, err := Connection()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec(stmt, uuid)
	if err != nil {
		return err
	}
	return nil
}
