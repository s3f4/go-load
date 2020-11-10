package library

import (
	"net/http"

	"github.com/gorilla/securecookie"
	"github.com/s3f4/mu/log"
)

var hashKey, _ = random(32)
var blockKey, _ = random(32)

var s = securecookie.New(hashKey, blockKey)

// SetCookie ...
func SetCookie(w http.ResponseWriter, c *http.Cookie, values map[string]string) error {
	encoded, err := s.Encode(c.Name, values)
	if err != nil {
		return err
	}

	cookie := &http.Cookie{
		Name:     c.Name,
		Value:    encoded,
		Secure:   c.Secure,
		HttpOnly: c.HttpOnly,
		Expires:  c.Expires,
	}

	if c.Path != "" {
		cookie.Path = c.Path
	} else {
		cookie.Path = "/"
	}

	if c.Domain != "" {
		cookie.Domain = c.Domain
	}

	http.SetCookie(w, cookie)
	return nil

}

// GetCookie ...
func GetCookie(r *http.Request, key string) (map[string]string, error) {
	cookie, err := r.Cookie(key)
	if err != nil {
		log.Debug("GetCookie value error:", err)
		return nil, err
	}

	value := make(map[string]string)
	if err = s.Decode(key, cookie.Value, &value); err != nil {
		log.Debug("GetCookie s.Decode error:", err)
		return nil, err
	}

	return value, nil
}
