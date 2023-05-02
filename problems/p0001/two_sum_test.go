package p0001

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Equal tells whether a and b contain the same elements.
// A nil argument is equivalent to an empty slice.
func Equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func TestTwoSum_Empty(t *testing.T) {
	// action
	var actual []int = twoSum([]int{}, 10)

	// asserts
	assert.True(t, Equal(actual, nil), "Expected empty slice")
}

func TestTwoSum_ResultDoesNotExist(t *testing.T) {
	// action
	var actual []int = twoSum([]int{1, 2}, 1)

	// asserts
	assert.True(t, Equal(actual, nil), "Expected empty slice")
}

func TestTwoSum_FindsResult(t *testing.T) {
	// action
	var actual []int = twoSum([]int{2, 7, 11, 15}, 9)

	// asserts
	assert.True(t, Equal(actual, []int{0, 1}), "Expected []int{0, 1}")
}
