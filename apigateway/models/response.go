package models

import "time"

// Response model
type Response struct {
	ID            uint
	TestID        uint
	TotalTime     int64
	FirstByte     time.Time
	FirstByteTime int64
	DNSStart      time.Time
	DNSDone       time.Time
	DNSTime       int64
	TLSStart      time.Time
	TLSDone       time.Time
	TLSTime       int64
	ConnectStart  time.Time
	ConnectDone   time.Time
	ConnectTime   int64
	StatusCode    int
	Body          string
}
