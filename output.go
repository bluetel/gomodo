package gomodo

import (
	"bytes"
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
	// Writes to the console, through various means
	Output()
}

// Standard OS output
type StreamOutput struct {
	w      io.Writer
	Buffer bytes.Buffer
}

// Creates a new stream output using os output
func NewStreamOutput() *StreamOutput {
	return &StreamOutput{
		w:      os.Stdout,
		Buffer: bytes.Buffer{},
	}
}

// Writes a string line
func (s *StreamOutput) Write(out string) {
	s.Buffer.WriteString(out)
}

// Writes a new line
func (s *StreamOutput) WriteLn(out string) {
	s.Buffer.WriteString(fmt.Sprintf("%s\n", out))
}

// Writes the buffered string to the stream
func (s *StreamOutput) Output() {
	s.Buffer.WriteTo(s.w)
}
