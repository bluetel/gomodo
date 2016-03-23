package gomodo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestApplicationCreation(t *testing.T) {
	app := NewApplication("Test", "v1.0.0")

	assert.Equal(t, "Test", app.Name, "The Application Name should be set")
	assert.Equal(t, "v1.0.0", app.Version, "The Application Version should be set")
}
