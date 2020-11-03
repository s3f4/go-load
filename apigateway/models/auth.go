package models

// TokenDetails model
type TokenDetails struct {
	AccessToken         string
	RefreshToken        string
	AccessUUID          string
	RefreshUUID         string
	AccessTokenExpires  int64
	RefreshTokenExpires int64
}

// AccessDetails model
type AccessDetails struct {
	AccessUUID string
	UserID     uint
}
