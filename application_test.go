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

// I Should be able to add an action to the Router
// And the router should be able to tell me this exists
func TestAddCommandToRouter(t *testing.T) {
	app := NewApplication("Test", "v1.0.0")

	resource := &ListTask{}
	app.AddResource(resource)

	_, success := app.Router.Find("list")
	assert.Equal(t, true, success, "The command should be found by the Router")
}

// I Should be able to run the application
// With no commands
func TestRunApplication(t *testing.T) {
	app := NewApplication("Test", "v1.0.0")
	app.Run()
}
