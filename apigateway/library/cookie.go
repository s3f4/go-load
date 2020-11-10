package library

import (
	"net/http"
	"os"

	"github.com/gorilla/securecookie"
)

var hashKey = []byte(os.Getenv("COOKIE_HASH_KEY"))
var blockKey = []byte(os.Getenv("COOKIE_BLOCK_KEY"))

var s = securecookie.New(hashKey, blockKey)

// SetCookie ...
func SetCookie(w http.ResponseWriter, c *http.Cookie, values map[string]string) {
	if encoded, err := s.Encode(c.Name, values); err == nil {
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
	}
}

// GetCookie ...
func GetCookie(r *http.Request, key string) map[string]string {
	if cookie, err := r.Cookie(key); err == nil {
		value := make(map[string]string)
		if err = s.Decode(key, cookie.Value, &value); err == nil {
			return value
		}
	}
	return nil
}
