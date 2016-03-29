package gomodo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// I should be able to create a new Router
// And add a route to it
// This should then be added to the Route Map
func TestCanAddRoutes(t *testing.T) {
	router := NewRouter()
	router.Add("test", "test")
	assert.Equal(t, 1, len(router.Routes))
}
