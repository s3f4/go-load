package models

// InstanceRequest is used for handlers and request
type InstanceRequest struct {
	InstanceCount    int    `json:"instanceCount"`
	InstanceSize     string `json:"size"`
	InstanceOS       string `json:"os"`
	Region           string `json:"region"`
	MaxWorkingPeriod int    `json:"maxWorkingPeriod"`
}
