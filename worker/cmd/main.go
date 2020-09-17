package main

import (
	"github.com/s3f4/go-load/worker/client"
)

func main() {
	q := client.NewRabbitMQService()

	// message := "hello there !"
	// if err := q.Send("worker", []byte(message)); err != nil {
	// 	log.Fatal(err)
	// }
	q.Listen("worker")
}
