package eventhandler

import (
	"fmt"
	"log"
	"os"

	"github.com/streadway/amqp"
)

// QueueService is a Queue Client service
// that connects to RabbitMQ takes messages and process
// those messages.
type QueueService interface {
	Send(queue string, message interface{}) error
	Listen(queue string)
}

type rabbitMQService struct {
	uri string
}

// NewRabbitMQService creates a new rabbitMQService instance
func NewRabbitMQService() QueueService {
	uri := fmt.Sprintf("amqp://%s:%s@%s:%s/",
		os.Getenv("RABBITMQ_USER"),
		os.Getenv("RABBITMQ_PASSWORD"),
		os.Getenv("RABBITMQ_HOST"),
		os.Getenv("RABBITMQ_PORT"),
	)

	return &rabbitMQService{uri: uri}
}

func (r *rabbitMQService) Send(queue string, message interface{}) error {
	conn, err := amqp.Dial(r.uri)
	if err != nil {
		return err
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		queue, // name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)

	if err != nil {
		return err
	}

	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/json",
			Body:        message.([]byte),
		},
	)

	return err
}

func (r *rabbitMQService) Listen(queue string) {
	conn, err := amqp.Dial(r.uri)
	if err != nil {
		log.Fatalf("%v failed to connect queue", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	defer ch.Close()

	msgs, err := ch.Consume(
		queue, // queue
		"",    // consumer
		false, // auto-ack
		false, // exclusive
		false, // no-local
		false, // no-wait
		nil,   // args
	)

	if err != nil {
		log.Fatalf("%v failed to register a consumer", err)
	}

	log.Printf("listening: %s", queue)

	block := make(chan struct{})
	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			ch.Ack(d.DeliveryTag, d.Redelivered)
		}
	}()
	fmt.Println("finishing...")
	<-block
}
