package client

import (
	"crypto/tls"
	"io/ioutil"
	"net/http"
	"net/http/httptrace"
	"time"

	"github.com/s3f4/go-load/worker/models"
	"github.com/s3f4/mu"
	"github.com/s3f4/mu/log"
)

// Client is a HTTP client that will be used for sending
// HTTP requests.
type Client struct {
	RunTestID       uint
	WorkerName      string
	URL             string
	TransportConfig models.TransportConfig
	Headers         []*models.Header
}

func (c *Client) setTrace(res *models.Response) *httptrace.ClientTrace {
	return &httptrace.ClientTrace{
		DNSStart: func(dsi httptrace.DNSStartInfo) {
			res.DNSStart = time.Now()
		},
		DNSDone: func(ddi httptrace.DNSDoneInfo) {
			res.DNSDone = time.Now()
		},
		TLSHandshakeStart: func() {
			res.TLSStart = time.Now()
		},
		TLSHandshakeDone: func(cs tls.ConnectionState, err error) {
			res.TLSDone = time.Now()
		},
		ConnectStart: func(network, addr string) {
			res.ConnectStart = time.Now()
		},
		ConnectDone: func(network, addr string, err error) {
			res.ConnectDone = time.Now()
		},
		GotFirstResponseByte: func() {
			res.FirstByte = time.Now()
		},
	}
}

func (c *Client) calculateTimes(start time.Time, res *models.Response) {
	res.FirstByteTime = int64(res.FirstByte.Sub(start))
	res.DNSTime = int64(res.DNSDone.Sub(res.DNSStart))
	res.TLSTime = int64(res.TLSDone.Sub(res.TLSStart))
	res.ConnectTime = int64(res.ConnectDone.Sub(res.ConnectStart))
	res.RunTestID = c.RunTestID
	res.WorkerHostName = getHostname()
	res.TotalTime = int64(time.Since(start))
}

// HTTPTrace load testing with HTTPTrace tool of golang.
func (c *Client) HTTPTrace() (*models.Response, error) {
	req, err := http.NewRequest("GET", c.URL, nil)

	if err != nil {
		log.Errorf("HTTPTrace Error: %v\n", err)
		return nil, err
	}

	log.Debugf("%#v\n", c)

	var res models.Response
	var start time.Time

	transport := http.DefaultTransport.(*http.Transport)
	transport.DisableKeepAlives = c.TransportConfig.DisableKeepAlives

	trace := c.setTrace(&res)

	req = req.WithContext(httptrace.WithClientTrace(req.Context(), trace))
	for _, header := range c.Headers {
		req.Header.Add(header.Key, header.Value)
	}

	start = time.Now()
	response, err := transport.RoundTrip(req)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	defer response.Body.Close()

	c.calculateTimes(start, &res)
	res.StatusCode = response.StatusCode

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	res.Body = string(body)
	return &res, err
}

// getHostname is used for finding out which worker making this requests.
func getHostname() string {
	output, _ := mu.RunCommands("hostname")
	return string(output)
}
