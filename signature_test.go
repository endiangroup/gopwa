package gopwa

import (
	"encoding/base64"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ItEscapesCharactersInStringToSign(t *testing.T) {
	method := "GET"
	ap := New(
		"abc123",
		"bcd234",
		"a1b2c3d4e5",
		UK,
		Sandbox,
	)
	queryParams := url.Values{
		"Action":         []string{"~DescribeJobFlows~"},
		"Version":        []string{"*2009-03-31*"},
		"AWSAccessKeyId": []string{"AKIAIOSFODNN7+EXAMPLE+"},
		"A+B":            []string{""},
		"A*B":            []string{""},
		"A~B":            []string{""},
	}

	toBeSigned := ap.prepareSignature(method, queryParams)

	assert.NotContains(t, toBeSigned, "+")
	assert.NotContains(t, toBeSigned, "*")
	assert.NotContains(t, toBeSigned, "~")

	assert.Contains(t, toBeSigned, "AKIAIOSFODNN7%20EXAMPLE%20")
	assert.Contains(t, toBeSigned, "%2A2009-03-31%2A")
	assert.Contains(t, toBeSigned, "%7EDescribeJobFlows%7E")
	assert.Contains(t, toBeSigned, "A%20B")
	assert.Contains(t, toBeSigned, "A%2AB")
	assert.Contains(t, toBeSigned, "A%7EB")
}

// See: http://docs.aws.amazon.com/general/latest/gr/signature-version-2.html
func Test_ItGeneratesExpectedHmacSHA256(t *testing.T) {
	method := "GET"
	ap := New(
		"",
		"",
		"wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY",
		UK,
		"",
	)
	ap.Endpoint, _ = url.Parse("https://elasticmapreduce.amazonaws.com/")

	queryParams := url.Values{
		"SignatureMethod":  []string{"HmacSHA256"},
		"SignatureVersion": []string{"2"},
		"Action":           []string{"DescribeJobFlows"},
		"Timestamp":        []string{"2011-10-03T15:19:30"},
		"Version":          []string{"2009-03-31"},
		"AWSAccessKeyId":   []string{"AKIAIOSFODNN7EXAMPLE"},
	}

	toBeSigned := ap.prepareSignature(method, queryParams)

	rawExpectedSignature, err := base64.StdEncoding.DecodeString("i91nKc4PWAt0JJIdXwz9HxZCJDdiy6cf/Mj6vPxyYIs=")
	assert.NoError(t, err)

	assert.Equal(t, rawExpectedSignature, ap.Signatory.Sign(toBeSigned))
}
