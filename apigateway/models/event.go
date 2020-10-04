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
	Event   EventType `json:"event"`
	Payload interface{} `json:"payload"`
}

// RequestPayload paylaod of request event
type RequestPayload struct {
	URL                  string          `json:"url"`
	Method               string          `json:"method"`
	Payload              string          `json:"payload,omitempty"`
	RequestCount         int             `json:"requestCount"`
	GoroutineCount       int             `json:"goroutineCount"`
	ExpectedResponseCode uint            `json:"expectedResponseCode"`
	ExpectedResponseBody string          `json:"expectedResponseBody"`
	TransportConfig      TransportConfig `json:"transportConfig"`
}
