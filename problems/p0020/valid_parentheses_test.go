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
