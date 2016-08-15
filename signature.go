package gopwa

import (
	"crypto/hmac"
	"crypto/sha256"
	"net/url"
	"strings"
)

var escapeCharactersTo = map[string]string{
	// Replaces encoded '+' (%2B) with encoded ' ' (%20)
	"%2B": "%20",
	"~":   "%7E",
}

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

// See: http://docs.aws.amazon.com/general/latest/gr/signature-version-2.html
func (ap AmazonPayments) prepareSignature(method string, queryParams url.Values) string {
	return strings.Join([]string{
		method,
		ap.Endpoint.Host,
		ap.Endpoint.Path,
		escapeCharacters(queryParams.Encode()),
	}, "\n")
}

func escapeCharacters(s string) string {
	for target, substitution := range escapeCharactersTo {
		s = strings.Replace(s, target, substitution, -1)
	}

	return s
}
