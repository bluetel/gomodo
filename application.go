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
	app.Arguments = os.Args
}

// Adds a resource to the Application
func (app *Application) AddResource(r Resource) {
	app.Router.AddResource(r)
}
