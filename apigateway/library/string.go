package library

import "crypto/rand"

const chars = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func random(length int) ([]byte, error) {
	bytes := make([]byte, length)

	if _, err := rand.Read(bytes); err != nil {
		return nil, err
	}

	for i, b := range bytes {
		bytes[i] = chars[b%byte(len(chars))]
	}

	return bytes, nil
}
