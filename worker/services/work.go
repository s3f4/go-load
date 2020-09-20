package services

import (
	"encoding/json"
	"strconv"

	"github.com/s3f4/go-load/worker/client"
	"github.com/s3f4/go-load/worker/models"
	"github.com/s3f4/mu/log"
)

// WorkService makes the load testing job.
type WorkService interface {
	Start(config *models.Work) error
}

type workerService struct {
	qs QueueService
}

// NewWorkService returns new workerService instance
func NewWorkService() WorkService {
	return &workerService{
		qs: NewRabbitMQService(),
	}
}

func (s *workerService) Start(config *models.Work) error {
	i := 0
	for i < config.GoroutineCount {
		log.Info("%+v", config)
		go s.run(config.URL, "worker_"+strconv.Itoa(i), config.Request)
		i++
	}
	return nil
}

func (s *workerService) run(url, workerName string, request int) {
	dataBuf := make(chan models.Response, 100)
	defer close(dataBuf)
	client := client.NewClient(
		url,
		workerName,
	)
	go s.makeReq(client, request, dataBuf)
	s.sendToEventHandler(dataBuf)
}

func (s *workerService) makeReq(client *client.Client, request int, dataBuf chan<- models.Response) {
	for i := 0; i < request; i++ {
		res := client.HTTPTrace()
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
