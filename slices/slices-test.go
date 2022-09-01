package slices

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_RemoveDuplicates(t *testing.T) {
	test := []string{"A", "B", "C", "D"}

	result := RemoveDuplicates(test)
	assert.Len(t, result, 4)

	test = append(test, "B")
	result = RemoveDuplicates(test)
	assert.Len(t, result, 4)
}
