package models

// Work model
type Work struct {
	Request         int             `json:"request"`
	URL             string          `json:"url"`
	GoroutineCount  int             `json:"goroutineCount"`
	TransportConfig TransportConfig `json:"transportConfig"`
}
