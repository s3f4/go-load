package models

// Response model
type Response struct {
	Request        int    `json:"request"`
	URL            string `json:"url"`
	GoroutineCount int    `json:"goroutineCount"`
}
