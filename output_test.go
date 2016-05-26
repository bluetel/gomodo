package gomodo

import (
	"bytes"
	_ "github.com/stretchr/testify/assert"
	"testing"
)

// We should be able to print a string
func TestItShouldBeAbleToWriteAString(t *testing.T) {
	stream := &StreamOutput{
		w: &bytes.Buffer{},
	}

	stream.Write("test")
}
