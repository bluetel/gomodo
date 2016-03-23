package main

import (
	"os"
)

// Application Object to create
// base CLI layer
type Application struct {
	// The Application Name
	Name string
	// The Application Version
	Version string
	// A list of Command Routes
	CommandRoutes []string
	// List of arguments
	Arguments []string
}

// Creates A new Application Struct and returns a pointer to this.
func NewApplication(name string, version string) *Application {
	return &Application{
		Name:    name,
		Version: version,
	}
}

// Runs the Application
func (app *Application) run() {
	// Get the arguments
	app.Arguments = os.Args
	// Pass to the validator
	// Find Command
	// Run Command
}

// Adds a resource to the Application
// This is how you get a command into the App
func (app *Application) AddResource(r Resource) {
	// Needs to read the results of a reflection
	// On the given resource class
}
