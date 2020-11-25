package models

// EventType to specify events
type EventType uint16

const (
	// REQUEST event type
	REQUEST EventType = iota
	// STORE events is used to store responses on db of the eventhandler
	STORE
)

// Event is a model which is used to queue communication
type Event struct {
	Event   EventType   `json:"event"`
	Payload interface{} `json:"payload"`
}

// RequestPayload paylaod of request event
type RequestPayload struct {
	RunTestID            uint            `json:"run_test_id"`
	URL                  string          `json:"url"`
	Method               string          `json:"method"`
	Payload              string          `json:"payload,omitempty"`
	RequestCount         uint64          `json:"request_count"`
	GoroutineCount       uint8           `json:"goroutine_count"`
	ExpectedResponseCode uint            `json:"expected_response_code"`
	ExpectedResponseBody string          `json:"expected_response_body"`
	TransportConfig      TransportConfig `json:"transport_config"`
	Headers              []*Header       `json:"headers"`
}

// Header holds request headers
type Header struct {
	Key             string `json:"key"`
	Value           string `json:"value"`
	IsRequestHeader bool   `json:"is_request_header"`
}

// TransportConfig is used to specify how to make requests
type TransportConfig struct {
	DisableKeepAlives bool `json:"disable_keep_alives"`
}
