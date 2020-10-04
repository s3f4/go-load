package models

import "time"

// TestConfig model
type TestConfig struct {
	ID        uint `json:"id" gorm:"primary_key"`
	StartTime *time.Time
	EndTime   *time.Time
	Tests     []*Test
}

// Test config to make requests
type Test struct {
	URL             string          `json:"url"`
	Method          string          `json:"method"`
	Payload         string          `json:"payload,omitempty"`
	RequestCount    int             `json:"requestCount"`
	GoroutineCount  int             `json:"goroutineCount"`
	StartTime       *time.Time      `json:"startTime"`
	EndTime         *time.Time      `json:"endTime"`
	TransportConfig TransportConfig `json:"transportConfig"`
}

// TransportConfig is used to specify how to make requests
type TransportConfig struct {
	DisableKeepAlives bool `json:"DisableKeepAlives"`
}
