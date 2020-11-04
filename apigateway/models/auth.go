package models

// AccessToken model
type AccessToken struct {
	Token  string
	UUID   string
	Expire int64
}

// RefreshToken model
type RefreshToken struct {
	Token  string
	UUID   string
	Expire int64
}

// Details model
type Details struct {
	UUID   string
	UserID uint
}
