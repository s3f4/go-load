package models

import "time"

// RunTest tests
type RunTest struct {
	ID        uint `json:"id" gorm:"primaryKey"`
	TestID    uint 
	StartTime *time.Time
	EndTime   *time.Time
	Passed    bool `json:"passed"`
}
