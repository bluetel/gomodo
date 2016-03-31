package gomodo

// A Resource controls Actions that can be performed
// These are command controllers that contain
// Further actions or a main action.
type Resource interface {
	GetName() string
	GetDescription() string
	PerformAction(app *Application)
}

// Default Help Task
type HelpTask struct {
}

func (t *HelpTask) GetName() string {
	return "help"
}

func (t *HelpTask) GetDescription() string {
	return "Provides help for a specific command, or all commands"
}

func (t *HelpTask) PerformAction(app *Application) {
	// Nothing to do here *yet*
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
	// Get All Commands
	// List them out
}
