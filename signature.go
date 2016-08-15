package gopwa

import (
	"crypto/hmac"
	"crypto/sha256"
)

type Signatory interface {
	Sign(string) []byte
	Method() string
	Version() string
}

type V2Hmac256Signatory struct {
	secret string
}

func (h V2Hmac256Signatory) Sign(raw string) []byte {
	hash := hmac.New(sha256.New, []byte(h.secret))

	hash.Write([]byte(raw))

	return hash.Sum(nil)
}

func (h V2Hmac256Signatory) Method() string {
	return "HmacSHA256"
}
func (h V2Hmac256Signatory) Version() string {
	return "2"
}
