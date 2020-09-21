package client

import (
	"crypto/tls"
	"io/ioutil"
	"net/http"
	"net/http/httptrace"
	"time"

	"github.com/s3f4/go-load/worker/models"
	"github.com/s3f4/mu/log"
)

// Client is a HTTP client that will be used for sending
// HTTP requests.
type Client struct {
	WorkerName string
	URL        string
	tc         models.TransportConfig
}

// HTTPTrace load testing with HTTPTrace tool of golang.
func (c *Client) HTTPTrace() *models.Response {
	req, err := http.NewRequest("GET", c.URL, nil)

	if err != nil {
		log.Errorf("HTTPTrace Error: %v\n", err)
		return nil
	}

	log.Infof("%#v\n", c)

	var res models.Response
	var start time.Time

	transport := http.DefaultTransport.(*http.Transport)
	transport.TLSHandshakeTimeout = time.Duration(c.tc.TLSHandshakeTimeout) * time.Second

	trace := &httptrace.ClientTrace{
		DNSStart:             func(dsi httptrace.DNSStartInfo) { res.DNSStart = time.Now().UTC() },
		DNSDone:              func(ddi httptrace.DNSDoneInfo) { res.DNSDone = time.Now() },
		TLSHandshakeStart:    func() { res.TLSStart = time.Now() },
		TLSHandshakeDone:     func(cs tls.ConnectionState, err error) { res.TLSDone = time.Now() },
		ConnectStart:         func(network, addr string) { res.ConnectStart = time.Now() },
		ConnectDone:          func(network, addr string, err error) { res.ConnectDone = time.Now() },
		GotFirstResponseByte: func() { res.FirstByte = time.Now() },
	}

	req = req.WithContext(httptrace.WithClientTrace(req.Context(), trace))
	start = time.Now()
	response, err := transport.RoundTrip(req)
	defer response.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	res.TotalTime = int64(time.Since(start))
	res.StatusCode = response.StatusCode
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Panic(err)
		return nil
	}
	res.Body = string(body)
	return &res
}
