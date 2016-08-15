package gopwa

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ItDefaultsToSandboxIfEmptyEnvironmentPassed(t *testing.T) {
	ap := New("", "", "", UK, "")

	expectedPath := Sandbox + "/2013-01-01"
	assert.Equal(t, string(expectedPath), ap.Endpoint.Path)
}

