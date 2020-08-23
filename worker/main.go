package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func connect() {
	conn, err := amqp.Dial("amqp://user:password@queue:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")

	body := "Hello World!"
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	failOnError(err, "Failed to publish a message")
}

func main() {
	connect()
	fmt.Println("test")
	i := 0
	for i < 1 {
		fmt.Println("test")
		DoRequest()
		i++
	}
}

func DoRequest() {
	start := time.Now()
	result, err := http.Get("https://s3f4.com/")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Request Time: %v, Request Code: %v \n", time.Since(start), result.StatusCode)
}
