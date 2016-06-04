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
	app.Output.Write(app.Tag)
	// Print application details
	app.Output.WriteLn("NAME:")
	app.Output.WriteLn(fmt.Sprintf("%4s %s - %s", "", app.Name, app.Description))

	app.Output.WriteLn("VERSION:")
	app.Output.WriteLn(fmt.Sprintf("%4s %s", "", app.Version))

	// Usage
	app.Output.WriteLn("USAGE:")
	app.Output.WriteLn(fmt.Sprintf("%4s %s [commmand] [options] [arguments]", "", app.Base))

	// Get All Commands
	routes := app.Router.GetRouteMap()

	// List them out
	app.Output.WriteLn("COMMANDS:")

	for key, route := range routes {
		usage := fmt.Sprintf("%s %s [options] [arguments]", app.Base, key)
		app.Output.WriteLn(fmt.Sprintf("%s %s - %s", RightPad(key, 20, " "), RightPad(usage, 40, " "), route.GetDescription()))
	}
}
