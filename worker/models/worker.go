package models

// Worker model
type Worker struct {
	Request        int    `json:"request"`
	URL            string `json:"url"`
	GoroutineCount int    `json:"goroutineCount"`
}
