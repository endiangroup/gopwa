package gopwa

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ItFormatsSingleError(t *testing.T) {
	errorResponse := ErrorResponse{
		RequestID: "abc123",
		Errors: []Error{
			Error{
				Type:    "Type",
				Code:    "Code ",
				Message: " Message. ",
				Detail:  "Detail",
			},
		},
	}

	assert.EqualError(t, errorResponse, "gopwa: Error for request 'abc123' - Code: Message")
}

func Test_ItFormatsMultipleErrors(t *testing.T) {
	errorResponse := ErrorResponse{
		RequestID: "abc123",
		Errors: []Error{
			Error{
				Type:    "Type1",
				Code:    "Code1",
				Message: "Message1",
				Detail:  "Detail1",
			},
			Error{
				Type:    "Type2",
				Code:    "Code2",
				Message: "Message2",
				Detail:  "Detail2",
			},
			Error{
				Type:    "Type3",
				Code:    "Code3",
				Message: "Message3",
				Detail:  "Detail3",
			},
		},
	}

	assert.EqualError(t, errorResponse, "gopwa: Error for request 'abc123' - Code1: Message1, Code2: Message2, Code3: Message3")
}
