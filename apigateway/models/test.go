package models

// Test config to make requests
type Test struct {
	ID                   uint            `json:"id" gorm:"id;primaryKey;autoIncrement"`
	TestGroupID          uint            `json:"test_group_id"`
	URL                  string          `json:"url" gorm:"url"`
	Method               string          `json:"method" gorm:"method"`
	Payload              string          `json:"payload,omitempty" gorm:"payload"`
	RequestCount         uint64          `json:"request_count" gorm:"request_count"`
	GoroutineCount       uint8           `json:"goroutine_count" gorm:"goroutine_count"`
	ExpectedResponseCode uint            `json:"expected_response_code" gorm:"expected_response_code"`
	ExpectedResponseBody string          `json:"expected_response_body" gorm:"expected_response_body"`
	TransportConfig      TransportConfig `json:"transport_config"`
	TestGroup            *TestGroup      `json:"test_group"`
	RunTests             []*RunTest      `json:"run_tests" gorm:"foreignKey:TestID"`
	Headers              []*Header       `json:"headers" gorm:"foreignKey:TestID"`
}

// Header holds request headers
type Header struct {
	ID              uint   `json:"id" gorm:"primaryKey;column:id"`
	TestID          uint   `json:"test_id"`
	Key             string `json:"key"`
	Value           string `json:"value"`
	IsRequestHeader bool   `json:"is_request_header"`
}

// TransportConfig is used to specify how to make requests
type TransportConfig struct {
	TestID            uint `json:"test_id" gorm:"foreignKey:TestID;column:test_id"`
	DisableKeepAlives bool `json:"disable_keep_alives" gorm:"disable_keep_alives"`
}
