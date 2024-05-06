package main

import (
	"fmt"

	"github.com/streadway/amqp"
)

func Rabbit() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer conn.Close()

	fmt.Println("successfully connected to rabbitmq")

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer ch.Close()

	msgs, err := ch.Consume(
		"MailQueue",
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	forever := make(chan bool)
	go func() {
		for d := range msgs {
			fmt.Println("Received message: ", string(d.Body))
		}
	}()

	fmt.Println("Successfully connected to RabbitMQ")
	fmt.Println("waiting for messages")
	<-forever
}
