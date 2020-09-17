package models

import "time"

// RunConfig config to make requests
type RunConfig struct {
	URL            int `json:"URL"`
	RequestCount   int `json:"requestCount"`
	GoroutineCount int `json:"goroutineCount"`
	InstanceCount  int `json:"instanceCount"`
	startTime      *time.Time
	endTime        *time.Time
}
