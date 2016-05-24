package gomodo

import (
	"os"
)

// Application Object to create
// base CLI layer
type Application struct {
	Name      string
	Version   string
	Arguments []string
	Options   []string
	Router    Router
}

// Creates A new Application Struct and returns a pointer to this.
func NewApplication(name string, version string) *Application {
	return &Application{
		Name:    name,
		Version: version,
		Router:  NewRouter(),
	}
}

// Runs the Application
func (app *Application) Run() {
	// Get the command name from os
	action := os.Args[1]
	// Does action exist
	_, success := app.Router.Find(action)

	if success {
		// Parse params
		// Run command
	}

	// List commands as the specified was not found
	cmd := &ListTask{}
	cmd.PerformAction(app)
}

// Sorts Arguments and Options
func (app *Application) sortParameters() {
}

// Adds a resource to the Application
func (app *Application) AddResource(r Resource) {
	app.Router.AddResource(r)
}
