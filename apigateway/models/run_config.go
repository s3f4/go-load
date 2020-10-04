package models

import "time"

// RunConfig config to make requests
type RunConfig struct {
	URL             string          `json:"url"`
	Method          string          `json:"method"`
	Payload         string          `json:"payload,omitempty"`
	RequestCount    int             `json:"requestCount"`
	GoroutineCount  int             `json:"goroutineCount"`
	InstanceCount   int             `json:"instanceCount"`
	StartTime       *time.Time      `json:"startTime"`
	EndTime         *time.Time      `json:"endTime"`
	TransportConfig TransportConfig `json:"transportConfig"`
}

// TransportConfig is used to specify how to make requests
type TransportConfig struct {
	DisableKeepAlives bool `json:"DisableKeepAlives"`
}
