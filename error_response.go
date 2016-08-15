package gopwa

import (
	"fmt"
	"strings"
)

// For error format see: http://docs.developer.amazonservices.com/en_US/dev_guide/DG_ResponseFormat.html
// For possible error codes see: https://payments.amazon.co.uk/developer/documentation/apireference/201753060
type ErrorResponse struct {
	StatusCode int     `xml:"-"`
	Errors     []Error `xml:"Error"`
	RequestID  string
}

func (e ErrorResponse) Error() string {
	output := fmt.Sprintf(
		"gopwa: Error for request '%s' - ",
		strings.TrimSpace(e.RequestID),
	)

	errs := make([]string, len(e.Errors))
	for i, err := range e.Errors {
		errs[i] = fmt.Sprintf(
			"%s: %s",
			strings.TrimSpace(err.Code),
			strings.TrimSuffix(strings.TrimSpace(err.Message), "."),
		)
	}

	output += strings.Join(errs, ", ")

	return output
}

type Error struct {
	Type    string
	Code    string
	Message string
	Detail  string
}
