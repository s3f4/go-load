package services

import (
	"strconv"

	"github.com/s3f4/go-load/worker"
	"github.com/s3f4/go-load/worker/client"
	"github.com/s3f4/go-load/worker/models"
	"github.com/s3f4/mu/log"
)

// WorkerService makes the load testing job.
type WorkerService interface {
	Start(config *models.Worker) error
	Done() error
}

type workerService struct{}

// NewWorkerService returns new workerService instance
func NewWorkerService() WorkerService {
	return &workerService{}
}

func (s *workerService) Start(config *models.Worker) error {
	i := 0
	for i < config.GoroutineCount {
		log.Info("%+v", config)
		go s.Run(config.URL, "worker_"+strconv.Itoa(i), config.Request)
		i++
	}
	return nil
}

func (s *workerService) Run(url, workerName string, request int) {
	dataBuf := make(chan worker.Response, 100)
	client := client.NewClient(
		url,
		workerName,
	)
	s.makeReq(request, dataBuf)
}

func (s *workerService) makeReq(request int, databuf chan<- worker.Response) {
	for i := 0; i < request; i++ {
		res := client.HTTPTrace()
		dataBuf <- *res
	}
}

func (s *workerService) insertToDB(dataBuf chan<- worker.Response) error {

	return nil
}
