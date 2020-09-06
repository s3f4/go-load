package models

// Instance is used for handlers and request
type Instance struct {
	ID               uint   `json:"id" gorm:"primary_key"`
	InstanceCount    int    `json:"instanceCount"`
	InstanceSize     string `json:"size"`
	InstanceOS       string `json:"os"`
	Region           string `json:"region"`
	MaxWorkingPeriod int    `json:"maxWorkingPeriod"`
}
