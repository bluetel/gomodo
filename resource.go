package gomodo

// A Resource controls Actions that can be performed
// These are command controllers that contain
// Further actions or a main action.
type Resource interface {
	GetName() string
}

// Default Help Task
type HelpTask struct {
}

// Gets the name of the help task
func (t *HelpTask) GetName() string {
	return "help"
}

// @Command list
// @Description Lists available commands
func (t *HelpTask) ListCommand(app *Application) {
	// Lists commands
}
