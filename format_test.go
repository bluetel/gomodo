package gomodo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// Should be able to pad a string from the left
func TestLeftPad(t *testing.T) {
	str := LeftPad("test", 6, "-")

	assert.Equal(t, "--test", str)
}

func TestRightPad(t *testing.T) {
	str := RightPad("test", 6, "-")

	assert.Equal(t, "test--", str)
}
