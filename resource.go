package gomodo

import (
	"fmt"
)

// A Resource controls Actions that can be performed
// These are command controllers that contain
// Further actions or a main action.
type Resource interface {
	GetName() string
	GetDescription() string
	PerformAction(app *Application)
}

// List task
type ListTask struct {
}

func (t *ListTask) GetName() string {
	return "list"
}

func (t *ListTask) GetDescription() string {
	return "Lists commands available in this application"
}

func (t *ListTask) PerformAction(app *Application) {
	// If Tag is set print tag
	fmt.Print(app.Tag)
	// Print application details
	fmt.Println("NAME:")
	fmt.Println(fmt.Sprintf("%4s %s - %s", "", app.Name, app.Description))

	fmt.Println("VERSION:")
	fmt.Println(fmt.Sprintf("%4s %s", "", app.Version))

	// Usage
	fmt.Println("USAGE:")
	fmt.Println(fmt.Sprintf("%4s %s [commmand] [options] [arguments]", "", app.Base))

	// Get All Commands
	routes := app.Router.GetRouteMap()

	// List them out
	fmt.Println("COMMANDS:")

	for key, route := range routes {
		usage := fmt.Sprintf("%s %s [options] [arguments]", app.Base, key)
		fmt.Println(fmt.Sprintf("%s %10s %32s - %s", key, "", usage, route.GetDescription()))
	}
}
