package util

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCaptureOutput(t *testing.T) {
	output := CaptureOutput(func() {
		fmt.Print("TEST")
	})
	assert.NotNil(t, output)
	assert.Len(t, output, 4)

	testFunc := func() {
		log.Println("Testing output to stdout")
		os.Stderr.Write([]byte("Testing output to stderr"))
	}

	// Capture the output
	output = CaptureOutput(testFunc)

	// Restore original output
	stdout := os.Stdout
	stderr := os.Stderr
	defer func() {
		os.Stdout = stdout
		os.Stderr = stderr
		log.SetOutput(os.Stderr)
	}()

	assert.Contains(t, output, "Testing output to stdout")
	assert.Contains(t, output, "Testing output to stderr")
}
