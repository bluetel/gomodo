package gomodo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// A Test Resource
type TestResource struct {
}

// Get Name Method
func (t *TestResource) GetName() string {
	return "Test"
}

// I should be able to create a new Router
// And add a route to it
// This should then be added to the Route Map
func TestCanAddRoutes(t *testing.T) {
	router := NewRouter()
	router.AddResource(&TestResource{})
	assert.Equal(t, 1, len(router.Routes))
}

// I should be able to add a new route
// And then Locate the struct name it belongs
// Using the Find method
func TestCanFindAddedRoute(t *testing.T) {
	router := NewRouter()
	router.AddResource(&TestResource{})

	obj, exists := router.Find("Test")
	assert.Equal(t, "TestResource", obj.Name())
	assert.Equal(t, true, exists)
}

// I should be able to check whether a route exists
// Using the same Find Method
func TestReturnFalseForRoutesNotFound(t *testing.T) {
	router := NewRouter()

	_, exists := router.Find("fakeAction")
	assert.Equal(t, false, exists)
}
