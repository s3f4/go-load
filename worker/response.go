package worker

import "time"

// Response model
type Response struct {
	TotalTime    time.Time
	FirstByte    time.Time
	DNSStart     time.Time
	DNSDone      time.Time
	TLSStart     time.Time
	TLSDone      time.Time
	ConnectStart time.Time
	ConnectDone  time.Time
	StatusCode   int
}
