package util

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDoIf(t *testing.T) {
	result := CaptureOutput(func() { DoIf(true, func() { fmt.Println("Test") }) })

	assert.Equal(t, "Test\n", result, "Should match the string Test")
}
