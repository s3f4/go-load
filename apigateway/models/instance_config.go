package models

//InstanceConfig is used for create new instances
type InstanceConfig struct {
	ID      uint        `json:"id" gorm:"primary_key,column:id"`
	Configs []*Instance `json:"configs" gorm:"foreignKey:instance_config_id"`
}

// Instance is used for handlers and request
type Instance struct {
	ID               uint   `json:"id" gorm:"primary_key,column:id"`
	InstanceConfigID uint   `json:"instance_config_id" gorm:"instance_config_id"`
	Count            int    `json:"count" gorm:"count"`
	InstanceSize     string `json:"size" gorm:"size"`
	Image            string `json:"image" gorm:"image"`
	Region           string `json:"region" gorm:"region"`
}
