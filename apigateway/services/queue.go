package services

import (
	"fmt"
	"os"

	"github.com/s3f4/go-load/apigateway/library/log"

	"github.com/streadway/amqp"
)

type processFunc func(d *amqp.Delivery, exit chan<- struct{}) error

// ListenSpec ...
type ListenSpec struct {
	Consumer    string
	Queue       string
	ProcessFunc processFunc
}

// QueueService is a Queue Client service
// that connects to RabbitMQ takes messages and process
// those messages.
type QueueService interface {
	Send(queue string, message interface{}) error
	Listen(listenSpec *ListenSpec)
	Declare(queue string) error
	Delete(queue string) error
}

type rabbitMQService struct {
	conn *amqp.Connection
}

// NewRabbitMQService creates a new rabbitMQService instance
func NewRabbitMQService() QueueService {
	uri := fmt.Sprintf("amqp://%s:%s@%s:%s/",
		os.Getenv("RABBITMQ_USER"),
		os.Getenv("RABBITMQ_PASSWORD"),
		os.Getenv("RABBITMQ_HOST"),
		os.Getenv("RABBITMQ_PORT"),
	)

	conn, err := amqp.Dial(uri)
	if err != nil {
		log.Fatalf("%v failed to connect queue", err)
	}

	return &rabbitMQService{
		conn: conn,
	}
}

func (r *rabbitMQService) Send(queue string, message interface{}) error {
	ch, err := r.conn.Channel()
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

func (r *rabbitMQService) Listen(spec *ListenSpec) {
	ch, err := r.conn.Channel()
	defer ch.Close()

	msgs, err := ch.Consume(
		spec.Queue,    // queue
		spec.Consumer, // consumer
		false,         // auto-ack
		false,         // exclusive
		false,         // no-local
		false,         // no-wait
		nil,           // args
	)

	if err != nil {
		log.Fatalf("%v failed to register a consumer", err)
	}

	block := make(chan struct{})
	go func() {
		for d := range msgs {
			if err := spec.ProcessFunc(&d, block); err != nil {
				log.Error(err)
			}
			ch.Ack(d.DeliveryTag, d.Redelivered)
		}
		log.Info("exit.")
	}()
	<-block
}

func (r *rabbitMQService) Declare(queue string) error {
	ch, err := r.conn.Channel()
	if err != nil {
		log.Debugf("QueueDeclare: conn.Channel: %s", err)
		return err
	}
	defer ch.Close()

	_, err = ch.QueueDeclare(
		queue, // name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)

	if err != nil {
		log.Debugf("QueueDeclare: queue declare: %s", err)
		return err
	}

	return nil
}

func (r *rabbitMQService) Delete(queue string) error {
	ch, err := r.conn.Channel()
	if err != nil {
		log.Debugf("Delete: conn.Channel: %s", err)
		return err
	}
	defer ch.Close()

	messageCount, err := ch.QueueDelete(queue, false, false, false)
	if err != nil {
		log.Debugf("QueueDeclare: queue declare: %s", err)
		return err
	}

	if messageCount > 0 {
		log.Errorf("Message count: %d > 0", messageCount)
	}

	return nil
}
