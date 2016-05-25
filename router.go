package gomodo

// Router Interface
type Router interface {
	// Method to add a resource
	AddResource(r Resource)
	// Method to find a resource
	Find(action string) (Resource, bool)
	// Returns a map of routes
	GetRouteMap() map[string]Resource
}

type MapRouter struct {
	Routes map[string]Resource
}

// Creates a new `Ready to use` Router struct
func NewRouter() *MapRouter {
	return &MapRouter{
		Routes: make(map[string]Resource),
	}
}

// Adds an entry for the given resource
func (m *MapRouter) AddResource(r Resource) {
	m.Routes[r.GetName()] = r
}

// Returns the found name and boolean flag as to whether it was successful
func (m *MapRouter) Find(action string) (Resource, bool) {
	obj, exists := m.Routes[action]
	return obj, exists
}

// Simply returns the Routes property, since its already a map
func (m *MapRouter) GetRouteMap() map[string]Resource {
	return m.Routes
}
