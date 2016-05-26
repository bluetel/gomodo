package gomodo

import (
	"github.com/jessevdk/go-flags"
	"os"
)

// Application Object to create
// base CLI layer
type Application struct {
	Name        string
	Description string
	Tag         string
	Version     string
	Base        string
	Arguments   []string
	Options     []string
	Router      Router
	Output      Output
}

// Creates A new Application Struct and returns a pointer to this.
func NewApplication(name string, description string, version string) *Application {
	return &Application{
		Name:        name,
		Description: description,
		Version:     version,
		Output:      NewStreamOutput(),
		Router:      NewRouter(),
	}
}

// Runs the Application
func (app *Application) Run() {
	// Set Base for file name
	app.Base = os.Args[0]

	list := &ListTask{}

	// Add the standard tasks
	app.AddResource(list)

	if len(os.Args) > 1 {
		// Get the command name from os
		action := os.Args[1]

		if app.Forward(action) {
			// Prevent further execution
			return
		}
	}

	// List commands as the specified was not found
	app.Forward("list")
}

// Used to execute a command, Inside another command
// it can be used to forward the action
func (app *Application) Forward(action string) bool {
	cmd, success := app.Router.Find(action)

	if success {
		// Parse params
		flags.Parse(cmd)
		// Run command
		cmd.PerformAction(app)
		// Commit the output
		app.Output.Output()
	}

	return success
}

// Adds a resource to the Application
func (app *Application) AddResource(r Resource) {
	app.Router.AddResource(r)
}
