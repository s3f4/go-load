package models

// InstanceRequest is used for handlers and request
type InstanceRequest struct {
	InstanceCount    int    `json:"instanceCount"`
	InstanceSize     string `json:"size"`
	Region           string `json:"region"`
	MaxWorkingPeriod int    `json:"maxWorkingPeriod"`
}
