package gomodo

import (
	_ "github.com/stretchr/testify/assert"
	"testing"
)

func TestAddResourceToApplication(t *testing.T) {
	app := NewApplication("test", "v1.0.0")

	// It should add the resource to the Application
	app.AddResource(&HelpTask{})
}
