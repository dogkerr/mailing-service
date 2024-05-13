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
			fmt.Println("=====================================")
			fmt.Println("Received a message: ", string(d.Body))
			fmt.Println("=====================================")

			var body structs.Message

			err := json.Unmarshal(d.Body, &body)
			if err != nil {
				fmt.Println("Error while reading JSON body:", err)
				continue
			}

			if err := body.Validate(); err != nil {
				fmt.Println("Error validating message:", err)
				continue
			}

			utils.SendGomail(structs.TemplateType(body.TemplateType), body.Data, body.Subject, body.To...)
		}
	}()

	fmt.Println("Successfully connected to RabbitMQ")
	fmt.Println("Waiting for messages...")
	<-forever
}
