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
