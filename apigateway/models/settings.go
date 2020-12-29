package models

// Settings holds settinsg
type Settings struct {
	ID      uint   `json:"id" gorm:"primaryKey;column:id"`
	Setting string `json:"setting" gorm:"column:setting;type:varchar(100)"`
	Value   string `json:"value" gorm:"column:value;type:varchar(100)"`
}

// SettingsKey for settings table
type SettingsKey string

const (
	// SIGNUP settings key
	SIGNUP SettingsKey = "SIGNUP"
)
