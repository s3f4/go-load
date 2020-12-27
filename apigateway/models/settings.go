package models

// Settings holds settinsg
type Settings struct {
	ID    uint   `json:"id" gorm:"primaryKey;column:id"`
	Key   string `json:"key" gorm:"key"`
	Value string `json:"value" gorm:"value"`
}
