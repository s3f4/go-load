package models

// Worker model
type Worker struct {
	Request string `json:"string"`
	URL     string `json:"url"`
	GoroutineCount int 
}
