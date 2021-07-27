package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/streadway/amqp"
)

func Publisher(client Client) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	conn, err := amqp.Dial(os.Getenv("AMQP_URL"))
	if err != nil {
		log.Fatal("Error conn Rabbit MQ")
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}
	defer ch.Close()

	queue := os.Getenv("AMQP_QUEUE")
	p, err := ch.QueueDeclare(
		queue,
		false,
		false,
		false,
		false,
		nil,
	)

	fmt.Println(p)
	if err != nil {
		fmt.Println(err)
	}

	// attempt to publish a message to the queue!
	err = ch.Publish(
		"",
		queue,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        []byte("hello world"),
		},
	)
}
