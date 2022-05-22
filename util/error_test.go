package util

import (
	"errors"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsErrorBool(t *testing.T) {
	output := CaptureOutput(func() {
		IsErrorBool(errors.New("Test Error"))
	})

	assert.Greater(t, len(output), 0, "No information logged to STDOUT")
	assert.Equal(t, strings.Count(output, "\n"), 1, "Expected one line of log output")
	assert.Contains(t, output, "Test Error")
}

func TestIfErrorLog(t *testing.T) {
	output := CaptureOutput(func() {
		IfErrorLog(errors.New("Test Error"))
	})

	assert.Greater(t, len(output), 0, "No information logged to STDOUT")
	assert.Equal(t, strings.Count(output, "\n"), 1, "Expected one line of log output")
	assert.Contains(t, output, "Test Error")
}
