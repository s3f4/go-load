package services

import (
	"encoding/json"
	"strconv"

	"github.com/mitchellh/mapstructure"
	"github.com/s3f4/go-load/worker/client"
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
	cfg := &mapstructure.DecoderConfig{
		Metadata: nil,
		Result:   &payload,
		TagName:  "json",
	}

	decoder, err := mapstructure.NewDecoder(cfg)
	if err != nil {
		log.Errorf("mapstructrure.decode", err)
		return err
	}

	if err := decoder.Decode(event.Payload); err != nil {
		log.Errorf("worker.start", err)
		return err
	}

	i := uint8(0)
	for i < payload.GoroutineCount {
		log.Info("%+v", payload)
		go s.run(payload.URL, "worker_"+strconv.Itoa(int(i)), payload.RunTestID, payload.RequestCount, payload.TransportConfig.DisableKeepAlives, payload.Headers)
		i++
	}
	return nil
}

// run
func (s *workerService) run(
	url, workerName string,
	runTestID uint,
	request uint64,
	disableKeepAlives bool,
	headers []*models.Header,
) {
	dataBuf := make(chan models.Response, 100)
	defer close(dataBuf)
	client := &client.Client{
		RunTestID:  runTestID,
		URL:        url,
		WorkerName: workerName,
		Headers:    headers,
		TransportConfig: models.TransportConfig{
			DisableKeepAlives: disableKeepAlives,
		},
	}
	go s.makeReq(client, request, dataBuf)
	s.sendToEventHandler(dataBuf)
}

func (s *workerService) makeReq(client *client.Client, request uint64, dataBuf chan<- models.Response) {
	// todo request/goroutineCount
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
