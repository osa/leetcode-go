package p0020

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValid_Empty(t *testing.T) {
	// action
	var actual bool = isValid("")

	// asserts
	assert.Equal(t, true, actual, "Expected true")
}

func TestIsValid_NonOverlapping(t *testing.T) {
	// action
	var actual bool = isValid("[]{}()")

	// asserts
	assert.Equal(t, true, actual, "Expected true")
}

func TestIsValid_Overlapping(t *testing.T) {
	// action
	var actual bool = isValid("([]{})")

	// asserts
	assert.Equal(t, true, actual, "Expected true")
}

func TestIsValid_NonMatchedClosing(t *testing.T) {
	// action
	var actual bool = isValid("[]{})'")

	// asserts
	assert.Equal(t, false, actual, "Expected false")
}

func TestIsValid_NonMatchedOpening(t *testing.T) {
	// action
	var actual bool = isValid("[]{'")

	// asserts
	assert.Equal(t, false, actual, "Expected false")
}
