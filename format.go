package gomodo

import (
	"bytes"
)

// Utility functions for padding and such
func LeftPad(in string, iterations int, pad string) string {
	var buffer bytes.Buffer
	var i int

	// Pads a string with the specified pad var
	// for the specified iterations
	iterations = iterations - len(in)

	for i = 0; i < iterations; i++ {
		buffer.WriteString(pad)
	}

	// Write original string
	buffer.WriteString(in)

	return buffer.String()
}

// Pads the given string to the right with the desired string
func RightPad(in string, iterations int, pad string) string {
	var buffer bytes.Buffer
	var i int

	iterations = iterations - len(in)

	buffer.WriteString(in)

	for i = 0; i < iterations; i++ {
		buffer.WriteString(pad)
	}

	return buffer.String()
}
