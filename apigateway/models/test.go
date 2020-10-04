package models

import "time"

// Test model
type Test struct {
	ID         uint `json:"id" gorm:"primary_key"`
	StartTime  *time.Time
	EndTime    *time.Time
	RunConfigs []*RunConfig
}
