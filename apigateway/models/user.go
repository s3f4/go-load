package models

// User model
type User struct {
	ID       uint   `json:"id" gorm:"id;primaryKey;autoIncrement"`
	Email    string `json:"email" gorm:"email;type:varchar(100);uniqueIndex;not null"`
	Password string `json:"password" gorm:"password;type:varchar(100);not null"`
}
