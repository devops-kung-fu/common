package util

import (
	"errors"
	"strings"
	"testing"

	"github.com/gookit/color"
	"github.com/stretchr/testify/assert"
)

func TestPrintTabbed(t *testing.T) {
	output := CaptureOutput(func() {
		PrintTabbed("[TEST]")
	})

	assert.Contains(t, output, "\t", "Console output does not contain a tab character")
	assert.Greater(t, len(output), 0, "No information logged to STDOUT")
	assert.Equal(t, strings.Count(output, "\n"), 1, "Expected one line of log output")
	assert.Equal(t, "\t[TEST]\n", output)
}

func TestPrintTabbedf(t *testing.T) {
	output := CaptureOutput(func() {
		PrintTabbedf("%s%s\n", "[TEST]", "123")
	})

	assert.Contains(t, output, "\t", "Console output does not contain a tab character")
	assert.Greater(t, len(output), 0, "No information logged to STDOUT")
	assert.Equal(t, strings.Count(output, "\n"), 1, "Expected one line of log output")
	assert.Equal(t, "\t[TEST]123\n", output)
}

func TestPrintSuccess(t *testing.T) {
	output := CaptureOutput(func() {
		PrintSuccess("[TEST]")
	})

	assert.Greater(t, len(output), 0, "No information logged to STDOUT")
	assert.Equal(t, strings.Count(output, "\n"), 1, "Expected one line of log output")
	assert.Equal(t, "[TEST]\n", output)
}

func TestPrintSuccessf(t *testing.T) {
	output := CaptureOutput(func() {
		PrintSuccessf("%s%s\n", "[TEST]", "123")
	})

	assert.Greater(t, len(output), 0, "No information logged to STDOUT")
	assert.Equal(t, strings.Count(output, "\n"), 1, "Expected one line of log output")
	assert.Equal(t, "[TEST]123\n", output)
}

func TestPrintWarning(t *testing.T) {
	output := CaptureOutput(func() {
		PrintWarning("[TEST]")
	})

	assert.Greater(t, len(output), 0, "No information logged to STDOUT")
	assert.Equal(t, strings.Count(output, "\n"), 1, "Expected one line of log output")
	assert.Equal(t, "[TEST]\n", output)
}

func TestPrintWarningf(t *testing.T) {
	output := CaptureOutput(func() {
		PrintWarningf("%s%s\n", "[TEST]", "123")
	})

	assert.Greater(t, len(output), 0, "No information logged to STDOUT")
	assert.Equal(t, strings.Count(output, "\n"), 1, "Expected one line of log output")
	assert.Equal(t, "[TEST]123\n", output)
}

func TestPrintInfo(t *testing.T) {
	output := CaptureOutput(func() {
		PrintInfo("[TEST]")
	})

	assert.Greater(t, len(output), 0, "No information logged to STDOUT")
	assert.Equal(t, strings.Count(output, "\n"), 1, "Expected one line of log output")
	assert.Equal(t, "[TEST]\n", output)
}

func TestPrintInfof(t *testing.T) {
	output := CaptureOutput(func() {
		PrintInfof("%s%s\n", "[TEST]", "123")
	})

	assert.Greater(t, len(output), 0, "No information logged to STDOUT")
	assert.Equal(t, strings.Count(output, "\n"), 1, "Expected one line of log output")
	assert.Equal(t, "[TEST]123\n", output)
}

func TestPrintErr(t *testing.T) {
	output := CaptureOutput(func() {
		PrintErr(errors.New("Test Error"))
	})

	assert.Greater(t, len(output), 0, "No information logged to STDOUT")
	assert.Equal(t, strings.Count(output, "\n"), 1, "Expected two lines of log output")
	assert.Equal(t, "Test Error\n", output)
}

func TestPrintErrf(t *testing.T) {
	output := CaptureOutput(func() {
		PrintErrf("This is a test: %s\n", errors.New("Test Error"))
	})

	assert.Greater(t, len(output), 0, "No information logged to STDOUT")
	assert.Equal(t, strings.Count(output, "\n"), 1, "Expected one line of log output")
	assert.Equal(t, "This is a test: Test Error\n", output)
}

func Test_printIcon(t *testing.T) {
	output := CaptureOutput(func() {
		printIcon(color.Red)
	})
	assert.Equal(t, len(output), 0, "No information logged to STDOUT")
	assert.Equal(t, "", output)
}
