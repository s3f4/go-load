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
	Portion      string   `json:"portion"`
	RequestCount uint64   `json:"request_count"`
	RunTest      *RunTest `json:"run_test"`
	Test         *Test    `json:"test"`
}
