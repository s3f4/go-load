package models

// Test config to make requests
type Test struct {
	ID                   uint `json:"id" gorm:"primaryKey"`
	TestGroupID          uint
	URL                  string          `json:"url"`
	Method               string          `json:"method"`
	Payload              string          `json:"payload,omitempty"`
	RequestCount         int             `json:"requestCount"`
	GoroutineCount       int             `json:"goroutineCount"`
	ExpectedResponseCode uint            `json:"expectedResponseCode"`
	ExpectedResponseBody string          `json:"expectedResponseBody"`
	TransportConfig      TransportConfig `json:"transportConfig"`
	RunTests             []*RunTest      `gorm:"foreignKey:TestID"`
}

// TransportConfig is used to specify how to make requests
type TransportConfig struct {
	TestID            uint `json:"test_id" gorm:"foreignKey:TestID"`
	DisableKeepAlives bool `json:"DisableKeepAlives"`
}
