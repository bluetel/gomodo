package gomodo

import (
	"os"
	"reflect"
)

// Application Object to create
// base CLI layer
type Application struct {
	Name      string
	Version   string
	Arguments []string
	Router    *Router
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
// This is how you get a command into the App
func (app *Application) AddResource(r Resource) {
	// Get the Name of the Struct from reflect
	obj := reflect.TypeOf(r).Name()
	// Get the Command name from struct method
	action := r.GetName()

	// Add to router
	app.Router.Add(action, obj)
}
