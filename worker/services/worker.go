package services

import (
	"strconv"

	"github.com/s3f4/go-load/worker/client"
)

// WorkerService makes the load testing job.
type WorkerService interface {
	Start(config interface{}) error
	Done() error
}

type workerService struct{}

// NewWorkerService returns new workerService instance
func NewWorkerService() WorkerService {
	return &workerService{}
}

func (s *workerService) Start(config interface{}) error {
	goRoutineCount := 10
	i := 0
	for i <= goRoutineCount {
		go s.Run("https://s3f4.com", "worker_"+strconv.Itoa(i))
		i++
	}
	return nil
}

func (s *workerService) Run(url, workerName string) {
	client := client.NewClient(
		url,
		workerName,
	)
	client.HTTPTrace()
}

func (s *workerService) Done() error { return nil }
