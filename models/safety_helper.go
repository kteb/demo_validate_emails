package models

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"hash"
)

// newHMAC creates and returns a new HMAC object
func NewHMAC(key string) HMAC {
	h := hmac.New(sha256.New, []byte(key))
	return HMAC{
		hmac: h,
	}
}

// lhmac is a wrapper arround the crypto/hmac package to make
// it a bit easier to use it in our code
type HMAC struct {
	hmac hash.Hash
}

// hash will hash the provided input string using HMAC with the
// secret key provided when the HMAC object was created
func (h HMAC) hash(input string) string {
	h.hmac.Reset()
	h.hmac.Write([]byte(input))
	b := h.hmac.Sum(nil)
	return base64.URLEncoding.EncodeToString(b)
}
