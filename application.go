package gomodo

import (
	"github.com/jessevdk/go-flags"
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
	if len(os.Args) > 1 {
		// Get the command name from os
		action := os.Args[1]
		// Does action exist
		command, success := app.Router.Find(action)

		if success {
			// Parse params
			flags.Parse(command)
			// Run command
			command.PerformAction(app)
			// Stop further execution
			return
		}
	}

	// List commands as the specified was not found
	cmd := &ListTask{}
	cmd.PerformAction(app)
}

// Adds a resource to the Application
func (app *Application) AddResource(r Resource) {
	app.Router.AddResource(r)
}
