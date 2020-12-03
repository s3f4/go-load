package services

import (
	"encoding/json"

	"github.com/s3f4/go-load/worker/client"
	"github.com/s3f4/go-load/worker/library"
	"github.com/s3f4/go-load/worker/models"
	"github.com/s3f4/mu/log"
)

// WorkerService makes the load testing job.
type WorkerService interface {
	Start(config *models.Event) error
}

type workerService struct {
	qs QueueService
}

var workerServiceObj WorkerService

// NewWorkerService returns new workerService instance
func NewWorkerService() WorkerService {
	if workerServiceObj == nil {
		workerServiceObj = &workerService{
			qs: NewRabbitMQService(),
		}
	}
	return workerServiceObj
}

// start gets started making requests.
func (s *workerService) Start(event *models.Event) error {
	var payload models.RequestPayload

	if err := library.DecodeMap(event.Payload, &payload); err != nil {
		log.Errorf("worker.start", err)
		return err
	}

	i := uint8(0)
	for i < payload.Test.GoroutineCount {
		log.Info("%+v", payload)
		go s.run(&payload)
		i++
	}
	return nil
}

// run
func (s *workerService) run(payload *models.RequestPayload) {
	// dataBuf allows eventhandler to save response results.
	dataBuf := make(chan models.Response, 100)
	defer close(dataBuf)

	client := &client.Client{
		RunTestID: payload.RunTest.ID,
		URL:       payload.Test.URL,
		Headers:   payload.Test.Headers,
		Method:    payload.Test.Method,
		TransportConfig: models.TransportConfig{
			DisableKeepAlives: payload.Test.TransportConfig.DisableKeepAlives,
		},
	}

	go s.makeReq(client, payload, dataBuf)
	s.sendToEventHandler(dataBuf)
}

func (s *workerService) makeReq(client *client.Client, payload *models.RequestPayload, dataBuf chan<- models.Response) {
	request := payload.RequestCount / uint64(payload.Test.GoroutineCount)
	for i := uint64(0); i < request; i++ {
		res, err := client.HTTPTrace()
		if err != nil {
			continue
		}
		dataBuf <- *res
	}
}

func (s *workerService) sendToEventHandler(dataBuf <-chan models.Response) error {
	for data := range dataBuf {
		jsonData, err := json.Marshal(data)
		if err != nil {
			log.Error(err)
			return err
		}
		s.qs.Send("eventhandler", jsonData)
	}

	return nil
}
