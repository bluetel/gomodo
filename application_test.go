package gomodo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// I Should be Able to Create a new Application
// And get the Name and Version
func TestApplicationCreation(t *testing.T) {
	app := NewApplication("Test", "v1.0.0")

	assert.Equal(t, "Test", app.Name, "The Application Name should be set")
	assert.Equal(t, "v1.0.0", app.Version, "The Application Version should be set")
}
