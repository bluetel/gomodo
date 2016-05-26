package gomodo

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

// We should be able to print a string
func TestItShouldBeAbleToWriteAString(t *testing.T) {
	stream := &StreamOutput{
		w: &bytes.Buffer{},
	}

	stream.Write("test")
	assert.Equal(t, "test", stream.Buffer.String(), "The test string should have been written to the buffer")
}
