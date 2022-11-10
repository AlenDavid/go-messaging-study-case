package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func main() {
	amqpUrl := flag.String("AMQP_URL", "amqp://guest:guest@my-stateful-broker-0.broker.default.svc.cluster.local:5672/", "The RabbitMQ url connection.")
	flag.Parse()

	conn, err := amqp.Dial(*amqpUrl)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var end chan struct{} = make(chan struct{})

	go func() {
		for {
			body := fmt.Sprintf("Hello World! at %s", time.Now())

			err = ch.PublishWithContext(ctx,
				"",     // exchange
				q.Name, // routing key
				false,  // mandatory
				false,  // immediate
				amqp.Publishing{
					Body: []byte(body),
				})

			if err != nil {
				log.Println("Failed to publish a message")
				end <- struct{}{}
			} else {
				log.Printf(" [x] Sent %s\n", body)
			}

			time.Sleep(5 * time.Second)
		}
	}()

	<-end
}
