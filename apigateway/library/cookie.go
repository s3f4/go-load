package library

import (
	"net/http"
	"os"

	"github.com/gorilla/securecookie"
)

var hashKey = []byte(os.Getenv("COOKIE_HASH_KEY"))
var blockKey = []byte(os.Getenv("COOKIE_BLOCK_KEY"))

var s = securecookie.New(hashKey, blockKey)

// SetCookieHandler ...
func SetCookieHandler(w http.ResponseWriter, key string, values map[string]string) {
	if encoded, err := s.Encode(key, values); err == nil {
		cookie := &http.Cookie{
			Name:     key,
			Value:    encoded,
			Path:     "/",
			Secure:   true,
			HttpOnly: true,
		}
		http.SetCookie(w, cookie)
	}
}

// ReadCookieHandler ...
func ReadCookieHandler(r *http.Request, key string) map[string]string {
	if cookie, err := r.Cookie(key); err == nil {
		value := make(map[string]string)
		if err = s.Decode(key, cookie.Value, &value); err == nil {
			return value
		}
	}
	return nil
}
