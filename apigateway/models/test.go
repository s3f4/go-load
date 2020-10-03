package models

import "time"

// Test model
type Test struct {
	ID        uint  `json:"id" gorm:"primary_key"`
	TestType  uint8 `json:"test_type"`
	StartTime *time.Time
	EndTime   *time.Time
}
