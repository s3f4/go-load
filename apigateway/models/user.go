package models

// User model
type User struct {
	ID       uint   `json:"id" gorm:"id,primaryKey"`
	Email    string `json:"email" gorm:"email"`
	Password string `json:"password" gorm:"password"`
}
