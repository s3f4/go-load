package models

import "time"

// Response model
type Response struct {
	ID              uint      `json:"id" gorm:"id;primaryKey;autoIncrement"`
	RunTestID       uint      `json:"run_test_id" gorm:"run_test_id"`
	WorkerHostName  string    `json:"worker_host_name" gorm:"worker_host_name"`
	TotalTime       int64     `json:"total_time" gorm:"total_time"`
	FirstByte       time.Time `json:"first_byte" gorm:"first_byte"`
	FirstByteTime   int64     `json:"first_byte_time" gorm:"first_byte_time"`
	DNSStart        time.Time `json:"dns_start" gorm:"dns_start"`
	DNSDone         time.Time `json:"dns_done" gorm:"dns_done"`
	DNSTime         int64     `json:"dns_time" gorm:"dns_time"`
	TLSStart        time.Time `json:"tls_start" gorm:"tls_start"`
	TLSDone         time.Time `json:"tls_done" gorm:"tls_done"`
	TLSTime         int64     `json:"tls_time" gorm:"tls_time"`
	ConnectStart    time.Time `json:"connect_start" gorm:"connect_start"`
	ConnectDone     time.Time `json:"connect_done" gorm:"connect_done"`
	ConnectTime     int64     `json:"connect_time" gorm:"connect_time"`
	StatusCode      int       `json:"status_code" gorm:"status_code"`
	Body            string    `json:"body" gorm:"body"`
	ResponseHeaders []byte    `json:"response_header" gorm:"response_header"`
	Error           string    `json:"error" gorm:"error"`
	Reasons         string    `json:"reasons" gorm:"reasons"`
	Passed          bool      `json:"passed" gorm:"passed"`
}
