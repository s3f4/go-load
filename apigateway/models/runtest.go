package models

import "time"

// RunTest tests
type RunTest struct {
	ID        uint       `json:"id" gorm:"primaryKey"`
	TestID    uint       `json:"test_id" gorm:"test_id"`
	StartTime *time.Time `json:"start_time" gorm:"start_time"`
	EndTime   *time.Time `json:"end_time" gorm:"end_time"`
	Passed    bool       `json:"passed" gorm:"passed"`
}
