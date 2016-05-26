package gomodo

import (
	"fmt"
	"io"
	"os"
)

// Output Interface to allow more control
// over what the application is printing
// and how that is formatted
type Output interface {
	// Method to write string to output
	Write(out string)
	// Method to write a new line
	WriteLn(out string)
}

// Standard OS output
type StreamOutput struct {
	w io.Writer
}

// Creates a new stream output using os output
func NewStreamOutput() *StreamOutput {
	return &StreamOutput{
		w: os.Stdout,
	}
}

// Writes a string line
func (s *StreamOutput) Write(out string) {
	io.WriteString(s.w, out)
}

// Writes a new line
func (s *StreamOutput) WriteLn(out string) {
	io.WriteString(s.w, fmt.Sprintf("%s\n", out))
}
