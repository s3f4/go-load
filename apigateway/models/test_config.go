package models

import "time"

// TestConfig model
type TestConfig struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	Name      string `json:"name"`
	StartTime *time.Time
	EndTime   *time.Time
	Tests     []*Test `json:"tests"`
}

// Test config to make requests
type Test struct {
	ID                   uint            `json:"id" gorm:"primaryKey"`
	TestConfigID         uint            `gorm:"foreignKey:TestConfigID"`
	URL                  string          `json:"url"`
	Method               string          `json:"method"`
	Payload              string          `json:"payload,omitempty"`
	RequestCount         int             `json:"requestCount"`
	GoroutineCount       int             `json:"goroutineCount"`
	ExpectedResponseCode uint            `json:"expectedResponseCode"`
	ExpectedResponseBody string          `json:"expectedResponseBody"`
	TransportConfig      TransportConfig `json:"transportConfig" gorm:"foreignKey:TestID"`
	Passed               bool            `json:"passed"`
	StartTime            *time.Time      `json:"startTime"`
	EndTime              *time.Time      `json:"endTime"`
}

// TransportConfig is used to specify how to make requests
type TransportConfig struct {
	TestID            uint `json:"test_id"`
	DisableKeepAlives bool `json:"DisableKeepAlives"`
}
