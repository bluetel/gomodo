package gomodo

import (
	"github.com/fatih/structs"
)

// Router Interface
type Router interface {
	// Method to add a resource
	AddResource(r Resource)
	// Method to find a resource
	Find(action string) (string, bool)
}

type MapRouter struct {
	Routes map[string]string
}

// Creates a new `Ready to use` Router struct
func NewRouter() *MapRouter {
	return &MapRouter{
		Routes: make(map[string]string),
	}
}

// Adds an entry for the given resource
func (m *MapRouter) AddResource(r Resource) {
	m.Routes[r.GetName()] = structs.Name(r)
}

// Returns the found name and boolean flag as to whether it was successful
func (m *MapRouter) Find(action string) (string, bool) {
	obj, exists := m.Routes[action]
	return obj, exists
}
