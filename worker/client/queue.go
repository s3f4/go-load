package client

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

// QueueService is a Queue Client service
// that connects to RabbitMQ takes messages and process
// those messages.
type QueueService interface {
	Send(queue string, message interface{}) error
	Receive(queue string) (interface{}, error)
	Listen(msgs <-chan amqp.Delivery, channel *amqp.Channel)
}

type rabbitMQService struct {
	uri string
}

// NewRabbitMQService creates a new rabbitMQService instance
func NewRabbitMQService(uri string) QueueService {
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
			ContentType: "text/plain",
			Body:        message.([]byte),
		},
	)

	return err

}
func (r *rabbitMQService) Receive(queue string) (message interface{}, err error) {
	conn, err := amqp.Dial(r.uri)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	ch, err := conn.Channel()
	defer ch.Close()

	q, err := ch.QueueDeclare(
		queue, // name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)

	msgs, err := ch.Consume(
		q.Name,   // queue
		"worker", // consumer
		false,    // auto-ack
		false,    // exclusive
		false,    // no-local
		false,    // no-wait
		nil,      // args
	)
	if err != nil {
		return nil, fmt.Errorf("%v failed to register a consumer", err)
	}

	r.Listen(msgs, ch)

	return nil, nil
}

func (r *rabbitMQService) Listen(msgs <-chan amqp.Delivery, ch *amqp.Channel) {
	go func() {
		for {
			for d := range msgs {
				log.Printf("Received a message: %s", d.Body)
				ch.Ack(d.DeliveryTag, d.Redelivered)
			}
		}
	}()
}
