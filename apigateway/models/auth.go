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

// AccessDetails model
type AccessDetails struct {
	AccessUUID string
	UserID     uint
}
