package models

//InstanceConfig is used for create new instances
type InstanceConfig struct {
	ID      uint        `json:"id" gorm:"primary_key"`
	Configs []*Instance `json:"configs"`
}

// Instance is used for handlers and request
type Instance struct {
	ID               uint   `json:"id" gorm:"primary_key"`
	InstanceCount    int    `json:"instanceCount"`
	InstanceSize     string `json:"size"`
	Image            string `json:"image"`
	Region           string `json:"region"`
	MaxWorkingPeriod int    `json:"maxWorkingPeriod"`
}
