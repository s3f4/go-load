package models

//InstanceConfig is used for create new instances
type InstanceConfig struct {
	ID      uint        `json:"id" gorm:"id;primaryKey;autoIncrement"`
	Configs []*Instance `json:"configs" gorm:"foreignKey:instance_config_id"`
}

// Instance is used for handlers and request
type Instance struct {
	ID               uint   `json:"id" gorm:"id;primaryKey;autoIncrement"`
	InstanceConfigID uint   `json:"instance_config_id" gorm:"instance_config_id"`
	Count            int    `json:"count" gorm:"count"`
	InstanceSize     string `json:"size" gorm:"size"`
	Image            string `json:"image" gorm:"image"`
	Region           string `json:"region" gorm:"region"`
}

// InstanceTerraform terraform configs
type InstanceTerraform struct {
	ID                 string `json:"id"`
	Name               string `json:"name"`
	CreatedAt          string `json:"created_at"`
	Disk               uint   `json:"disk"`
	Image              string `json:"image"`
	IPV4Address        string `json:"ipv4_address"`
	IPV4AddressPrivate string `json:"ipv4_address_private"`
	Memory             uint   `json:"memory"`
	Region             string `json:"region"`
	Size               string `json:"size"`
	Status             string `json:"status"`
}
