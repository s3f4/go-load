package models

import "time"

// Report is used to show request-response results
type Report struct {
	ReportType uint16 `json:"report_type" gorm:"report_type"`
	Data       []*Response 
	StartTime  *time.Time
	EndTime    *time.Time
}
