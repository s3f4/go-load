package models

import "time"

// RunConfig config to make requests
type RunConfig struct {
	URL            string     `json:"url"`
	RequestCount   int        `json:"requestCount"`
	GoroutineCount int        `json:"goroutineCount"`
	InstanceCount  int        `json:"instanceCount"`
	StartTime      *time.Time `json:"startTime"`
	EndTime        *time.Time `json:"endTime"`
}
