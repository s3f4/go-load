package models

import "time"

// Response model
type Response struct {
	TotalTime    int64
	FirstByte    time.Time
	DNSStart     time.Time
	DNSDone      time.Time
	TLSStart     time.Time
	TLSDone      time.Time
	ConnectStart time.Time
	ConnectDone  time.Time
	StatusCode   int
	Body         string
}
