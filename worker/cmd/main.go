package main

import (
	"log"

	"github.com/s3f4/go-load/worker/client"
)

func main() {
	queueURI := "amqp://user:password@queue:5672/"
	q := client.NewRabbitMQService(queueURI)

	message := "hello there !"
	if err := q.Send("worker", []byte(message)); err != nil {
		log.Fatal(err)
	}
	//q.Listen("test")
}
