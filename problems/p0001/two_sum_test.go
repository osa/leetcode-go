package p0001

import (
	"testing"
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
	if !Equal(actual, nil) {
		t.Errorf("Expected empty slice")
		t.Fatal()
	}
}

func TestTwoSum_ResultDoesNotExist(t *testing.T) {
	// action
	var actual []int = twoSum([]int{1, 2}, 1)

	// asserts
	if !Equal(actual, nil) {
		t.Errorf("Expected empty slice")
		t.Fatal()
	}
}

func TestTwoSum_FindsResult(t *testing.T) {
	// action
	var actual []int = twoSum([]int{2, 7, 11, 15}, 9)

	// asserts
	if !Equal(actual, []int{0, 1}) {
		t.Errorf("Expected []int{0, 1}")
		t.Fatal()
	}
}
