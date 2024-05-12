package main

import (
	"encoding/json"
	"fmt"

	"github.com/dogkerr/mailing-service/m/v2/structs"
	"github.com/dogkerr/mailing-service/m/v2/utils"
	"github.com/streadway/amqp"
)

func main() {
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
			var body structs.Message

			err := json.Unmarshal(d.Body, &body)
			if err != nil {
				fmt.Println("Error while reading JSON body")
				continue
			}

			utils.SendGomail(structs.TemplateType(body.TemplateType), body.Data, body.Subject, body.To...)
		}
	}()

	fmt.Println("Successfully connected to RabbitMQ")
	fmt.Println("waiting for messages")
	<-forever
}
