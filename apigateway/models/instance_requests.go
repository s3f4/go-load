package models

// InstanceRequest is used for handlers and request
type InstanceRequest struct {
	Count  int    `json:"count"`
	Region string `json:"region"`
}

