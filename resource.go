package gomodo

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
	// Get All Commands
	// List them out
}
