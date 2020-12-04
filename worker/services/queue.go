package services

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/s3f4/go-load/worker/library"
	"github.com/s3f4/go-load/worker/models"
	"github.com/s3f4/mu/log"
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
	conn *amqp.Connection
}

var rabbitMQServiceObject QueueService

// NewRabbitMQService creates a new rabbitMQService instance
func NewRabbitMQService() QueueService {
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
		}
	}
	return rabbitMQServiceObject
}

func (r *rabbitMQService) Send(queue string, message interface{}) error {
	ch, err := r.conn.Channel()
	if err != nil {
		log.Error(err)
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
		log.Error(err)
		return err
	}

	if err := ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/json",
			Body:        message.([]byte),
		},
	); err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func (r *rabbitMQService) Listen(queue string) {
	ch, err := r.conn.Channel()
	defer ch.Close()

	msgs, err := ch.Consume(
		queue,    // queue
		"worker", // consumer
		false,    // auto-ack
		false,    // exclusive
		false,    // no-local
		false,    // no-wait
		nil,      // args
	)

	if err != nil {
		log.Fatalf("%v failed to register a consumer", err)
	}

	block := make(chan struct{})
	go func() {
		for d := range msgs {
			time.Sleep(time.Second * 3)
			log.Infof("Received a message: %s", d.Body)

			var event models.Event
			if err := json.Unmarshal(d.Body, &event); err != nil {
				log.Errorf("worker json error: %s", err)
			}

			var payload models.RequestPayload

			if err := library.DecodeMap(event.Payload, &payload); err != nil {
				log.Errorf("worker decode_map error: %s", err)
			}

			s := NewWorkerService()
			s.Start(&payload)

			// Send latest workers done message
			portion := strings.Split(payload.Portion, "/")
			if portion[0] == portion[1] {
				q := fmt.Sprintf("collect_%d_%d", payload.Test.ID, payload.RunTest.ID)
				message, _ := json.Marshal(models.Event{
					Event: models.COLLECT,
					Payload: &models.CollectPayload{
						TestID:    payload.Test.ID,
						RunTestID: payload.RunTest.ID,
						Portion:   payload.Portion,
					},
				})
				r.Send(q, message)
			}
			// Done
			ch.Ack(d.DeliveryTag, d.Redelivered)
		}
	}()
	<-block
}
