package client

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"net/http/httptrace"
	"time"

	"github.com/s3f4/go-load/worker"
)

// Client is a HTTP client that will be used for sending
// HTTP requests.
type Client struct {
	workerName string
	url        string
}

// NewClient returns new Client instance
func NewClient(url, workerName string) *Client {
	return &Client{
		url:        url,
		workerName: workerName,
	}
}

// HTTPTrace load testing with HTTPTrace tool of golang.
func (c *Client) HTTPTrace() {
	req, err := http.NewRequest("GET", c.url, nil)
	if err != nil {
		log.Printf("HTTPTrace Error: %v\n", err)
		return
	}

	var res worker.Response
	var start time.Time

	trace := &httptrace.ClientTrace{
		DNSStart:             func(dsi httptrace.DNSStartInfo) { res.DNSStart = time.Now() },
		DNSDone:              func(ddi httptrace.DNSDoneInfo) { res.DNSDone = time.Now() },
		TLSHandshakeStart:    func() { res.TLSStart = time.Now() },
		TLSHandshakeDone:     func(cs tls.ConnectionState, err error) { res.TLSDone = time.Now() },
		ConnectStart:         func(network, addr string) { res.ConnectStart = time.Now() },
		ConnectDone:          func(network, addr string, err error) { res.ConnectDone = time.Now() },
		GotFirstResponseByte: func() { res.FirstByte = time.Now() },
	}

	req = req.WithContext(httptrace.WithClientTrace(req.Context(), trace))
	start = time.Now()
	if _, err := http.DefaultTransport.RoundTrip(req); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Total time: %v\n", time.Since(start))
}
