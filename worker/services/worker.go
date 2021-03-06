package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"sync"

	"github.com/s3f4/go-load/worker/client"
	"github.com/s3f4/go-load/worker/library"
	"github.com/s3f4/go-load/worker/models"
	"github.com/s3f4/mu/log"
)

// WorkerService makes the load testing job.
type WorkerService interface {
	Start(*models.RequestPayload)
}

type workerService struct {
	qs QueueService
}

// NewWorkerService returns new workerService instance
func NewWorkerService() WorkerService {
		return &workerService{
			qs: NewQueueService(),
		}
}

// start gets started making requests.
func (s *workerService) Start(payload *models.RequestPayload) {
	wg := &sync.WaitGroup{}
	i := uint8(0)
	for i < payload.Test.GoroutineCount {
		wg.Add(1)
		log.Info("%+v", payload)
		go s.run(payload, wg)
		i++
	}
	wg.Wait()
}

// run
func (s *workerService) run(payload *models.RequestPayload, wg *sync.WaitGroup) {
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

	go s.makeReq(client, payload, dataBuf, wg)
	s.sendToEventHandler(dataBuf)
}

func (s *workerService) makeReq(client *client.Client, payload *models.RequestPayload, dataBuf chan<- models.Response, wg *sync.WaitGroup) {
	request := payload.RequestCount / uint64(payload.Test.GoroutineCount)
	for i := uint64(0); i < request; i++ {
		res, err := client.HTTPTrace()
		if err != nil {
			res = &models.Response{}
			res.Error = err.Error()
		}

		reasons := s.compare(payload.Test, res)
		if len(reasons) > 0 {
			f := false
			res.Passed = &f
			res.Reasons = strings.Join(reasons, "\n")
			payload.RunTest.Passed = &f
		}
		dataBuf <- *res
		fmt.Printf("%#v\n", res.Passed)
	}
	wg.Done()
}

func (s *workerService) compare(test *models.Test, response *models.Response) []string {
	var reasons []string

	if test.ExpectedConnectionTime != 0 && test.ExpectedConnectionTime > response.ConnectTime {
		reasons = append(reasons, fmt.Sprintf(
			"test.ExpectedConnectionTime: %d > response.ConnectTime: %d",
			test.ExpectedConnectionTime,
			response.ConnectTime,
		))
	}

	if test.ExpectedTLSTime != 0 && test.ExpectedTLSTime > response.TLSTime {
		reasons = append(reasons, fmt.Sprintf(
			"test.ExpectedTLSTime: %d > response.TLSTime: %d",
			test.ExpectedTLSTime,
			response.TLSTime,
		))
	}

	if test.ExpectedDNSTime != 0 && test.ExpectedDNSTime > response.DNSTime {
		reasons = append(reasons, fmt.Sprintf(
			"test.ExpectedDNSTime: %d > response.DNSTime: %d",
			test.ExpectedDNSTime,
			response.DNSTime,
		))
	}

	if test.ExpectedFirstByteTime != 0 && test.ExpectedFirstByteTime > response.FirstByteTime {
		reasons = append(reasons, fmt.Sprintf(
			"test.ExpectedFirstByteTime: %d > response.FirstByteTime: %d",
			test.ExpectedFirstByteTime,
			response.FirstByteTime,
		))
	}

	var responseHeaders http.Header
	if err := json.Unmarshal(response.ResponseHeaders, &responseHeaders); err != nil {
		reasons = append(reasons, fmt.Sprintf(
			"ResponseHeaders json error: %v, Headers: %v",
			err,
			responseHeaders,
		))
	}

	for _, header := range test.Headers {
		if !header.IsRequestHeader {
			if values, ok := responseHeaders[header.Key]; ok {
				if found := library.SliceFind(values, header.Value); found == -1 {
					reasons = append(reasons, fmt.Sprintf(
						"header.Value: %s is not found in %v",
						header.Value,
						values,
					))
				}
			} else {
				reasons = append(reasons, fmt.Sprintf(
					"header: %s is not found in the response headers",
					header.Key,
				))
			}
		}
	}

	if test.ExpectedResponseCode != 0 && test.ExpectedResponseCode != response.StatusCode {
		reasons = append(reasons, fmt.Sprintf(
			"test.ExpectedResponseCode: %d, response.StatusCode: %d",
			test.ExpectedResponseCode,
			response.StatusCode,
		))
	}

	if test.ExpectedResponseBody != "" && test.ExpectedResponseBody != response.Body {
		reasons = append(reasons, fmt.Sprintf(
			"test.ExpectedResponseBody: %s, response.Body: %s",
			test.ExpectedResponseBody,
			response.Body,
		))
	}

	return reasons
}

func (s *workerService) sendToEventHandler(dataBuf <-chan models.Response) error {
	fmt.Println("databuf")
	fmt.Println(dataBuf)
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
