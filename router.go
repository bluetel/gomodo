package gomodo

type Router struct {
	Routes map[string]string
}

// Creates a new `Ready to use` Router struct
func NewRouter() *Router {
	return &Router{
		Routes: make(map[string]string),
	}
}

// Adds an entry for the route using Action -> Object
func (r *Router) Add(action string, obj string) {
	r.Routes[action] = obj
}

// Returns the found name and boolean flag as to whether it was successful
func (r *Router) Find(action string) (string, bool) {
	obj, exists := r.Routes[action]
	return obj, exists
}
