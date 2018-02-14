package actions

import (
	"crypto/rand"
	"encoding/base64"
)

const generateTokenBytes = 32

// generateToken is a helper function designed to generate
// tokens of a predetermined byte size
func generateToken() (string, error) {
	return stringToken(generateTokenBytes)
}

//stringToken will generate a byte slice of size nBytes and the return a string
func stringToken(nBytes int) (string, error) {
	b, err := bytes(nBytes)
	if (err) != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}

// bytes will help us generate n random bytes, or will
// return an error if there was one. This uses the crypto/rand
// package so it is safe to use with things like remember tokens
func bytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}

// nBytes returns the number of bytes used in the base64
// URL encoded string
func nBytes(base64String string) (int, error) {
	b, err := base64.URLEncoding.DecodeString(base64String)
	if err != nil {
		return -1, err
	}
	return len(b), err
}
