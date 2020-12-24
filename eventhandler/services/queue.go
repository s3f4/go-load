package services

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/s3f4/go-load/eventhandler/models"
	"github.com/s3f4/go-load/eventhandler/repository"
	"github.com/s3f4/mu/log"
	"gorm.io/gorm"

	"github.com/streadway/amqp"
)

// QueueService is a Queue Client service
// that connects to RabbitMQ takes messages and process
// those messages.
type QueueService interface {
	Send(queue string, message interface{}) error
	Listen(queue string)
	QueueDeclare(queue string) error
}

type rabbitMQService struct {
	db   *gorm.DB
	conn *amqp.Connection
}

var rabbitMQServiceObject QueueService

// NewRabbitMQService creates a new rabbitMQService instance
func NewRabbitMQService(db *gorm.DB) QueueService {
	if rabbitMQServiceObject == nil {
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

		rabbitMQServiceObject = &rabbitMQService{
			conn: conn,
			db:   db,
		}
	}
	return rabbitMQServiceObject
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

func (r *rabbitMQService) Listen(queue string) {
	ch, err := r.conn.Channel()
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

	block := make(chan struct{})
	responseRepository := repository.NewResponseRepository(r.db)
	go func() {
		for d := range msgs {
			log.Infof("Received a message: %s", d.Body)
			var resp models.Response
			if err := json.Unmarshal(d.Body, &resp); err != nil {
				log.Error(err)
			}
			responseRepository.Create(&resp)
			ch.Ack(d.DeliveryTag, d.Redelivered)
		}
	}()
	<-block
	log.Infof("finished listening the queue of %s")
}

func (r *rabbitMQService) QueueDeclare(queue string) error {
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
