package models

// InstanceRequest is used for handlers and request
type InstanceRequest struct {
	RequestCount     int    `json:"requestCount"`
	InstanceCount    int    `json:"instanceCount"`
	Region           string `json:"region"`
	MaxWorkingPeriod int    `json:"maxWorkingPeriod"`
}
