package services

import (
	"fmt"
	"os"

	"github.com/s3f4/go-load/worker/library/specs"
	"github.com/s3f4/mu/log"
	"github.com/streadway/amqp"
)

// QueueService is a Queue Client service
// that connects to RabbitMQ takes messages and process
// those messages.
type QueueService interface {
	Send(queue string, message interface{}) error
	Listen(spec *specs.ListenSpec)
	QueueDeclare(queue string) error
}

type queueService struct {
	conn *amqp.Connection
}

// NewQueueService creates a new queueService instance
func NewQueueService() QueueService {
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

	return &queueService{
		conn: conn,
	}
}

func (r *queueService) Send(queue string, message interface{}) error {
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

func (r *queueService) Listen(spec *specs.ListenSpec) {
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

func (r *queueService) QueueDeclare(queue string) error {
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
