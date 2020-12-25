package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/s3f4/go-load/worker/library"
	"github.com/s3f4/go-load/worker/library/specs"
	"github.com/s3f4/go-load/worker/models"
	"github.com/s3f4/go-load/worker/services"
	"github.com/s3f4/mu/log"
	"github.com/streadway/amqp"
)

func main() {
	queueService := services.NewQueueService()
	listener := services.NewListener()
	worker := services.NewWorkerService()

	spec := &specs.ListenSpec{
		Consumer: "worker",
		Queue:    "worker",
		ProcessFunc: func(d *amqp.Delivery, exit chan<- struct{}) error {

			log.Debugf("Message: %s", d.Body)
			var event models.Event
			if err := json.Unmarshal(d.Body, &event); err != nil {
				log.Errorf("worker json error: %s", err)
			}

			var payload models.RequestPayload
			if err := library.DecodeMap(event.Payload, &payload); err != nil {
				log.Errorf("worker decode_map error: %s", err)
			}

			worker.Start(&payload)

			// finishing test
			endTime := time.Now()
			payload.RunTest.EndTime = &endTime
			// Send latest workers done message
			q := fmt.Sprintf("collect_%d_%d", payload.Test.ID, payload.RunTest.ID)
			message, _ := json.Marshal(models.Event{
				Event: models.COLLECT,
				Payload: &models.CollectPayload{
					Test:    payload.Test,
					RunTest: payload.RunTest,
					Portion: payload.Portion,
				},
			})
			queueService.Send(q, message)
			return nil
		},
	}

	specs := []*specs.ListenSpec{
		spec,
	}

	listener.Start(specs)
}
