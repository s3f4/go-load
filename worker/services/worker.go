package services

import (
	"encoding/json"
	"strconv"

	"github.com/mitchellh/mapstructure"
	"github.com/s3f4/go-load/worker/client"
	"github.com/s3f4/go-load/worker/models"
	"github.com/s3f4/mu"
	"github.com/s3f4/mu/log"
)

// WorkService makes the load testing job.
type WorkService interface {
	Start(config *models.Event) error
}

type workerService struct {
	qs QueueService
}

// NewWorkerService returns new workerService instance
func NewWorkerService() WorkService {
	return &workerService{
		qs: NewRabbitMQService(),
	}
}

func (s *workerService) Start(event *models.Event) error {
	var payload models.RequestPayload
	mapstructure.Decode(event.Payload, &payload)

	i := 0
	for i < payload.GoroutineCount {
		log.Info("%+v", payload)
		go s.run(payload.URL, "worker_"+strconv.Itoa(i), payload.RequestCount, payload.TransportConfig.DisableKeepAlives)
		i++
	}
	return nil
}

func (s *workerService) run(url, workerName string, request int, disableKeepAlives bool) {
	dataBuf := make(chan models.Response, 100)
	defer close(dataBuf)
	client := &client.Client{
		URL:        url,
		WorkerName: workerName,
		TransportConfig: models.TransportConfig{
			DisableKeepAlives: disableKeepAlives,
		},
	}
	go s.makeReq(client, request, dataBuf)
	s.sendToEventHandler(dataBuf)
}

func (s *workerService) makeReq(client *client.Client, request int, dataBuf chan<- models.Response) {
	url := client.URL
	for i := 0; i < request; i++ {
		client.URL = url + "/?reqRef=" + mu.RandomString(5)
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
