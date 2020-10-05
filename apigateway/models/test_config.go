package models

import "time"

// TestConfig model
type TestConfig struct {
	ID        uint   `json:"id" gorm:"primary_key"`
	Name      string `json:"name"`
	StartTime *time.Time
	EndTime   *time.Time
	Tests     []*Test
}

// Test config to make requests
type Test struct {
	URL                  string          `json:"url"`
	Method               string          `json:"method"`
	Payload              string          `json:"payload,omitempty"`
	RequestCount         int             `json:"requestCount"`
	GoroutineCount       int             `json:"goroutineCount"`
	ExpectedResponseCode uint            `json:"expectedResponseCode"`
	ExpectedResponseBody string          `json:"expectedResponseBody"`
	TransportConfig      TransportConfig `json:"transportConfig"`
	Passed               bool            `json:"passed"`
	StartTime            *time.Time      `json:"startTime"`
	EndTime              *time.Time      `json:"endTime"`
}

// TransportConfig is used to specify how to make requests
type TransportConfig struct {
	DisableKeepAlives bool `json:"DisableKeepAlives"`
}
