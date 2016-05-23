package gomodo

import (
	"flag"
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
		// Parse arguments
		app.sortParameters()
		// Fire the command action
	}

	// List commands as the specified was not found
	cmd := &ListTask{}
	cmd.PerformAction(app)
}

// Sorts Arguments and Options
func (app *Application) sortParameters() {
	// Parse Flags
	flag.Parse()
	// Extract Arguments
	app.Arguments = flag.Args()
}

// Adds a resource to the Application
func (app *Application) AddResource(r Resource) {
	app.Router.AddResource(r)
}
