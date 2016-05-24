[![Build Status](https://travis-ci.org/Bluetel-Solutions/gomodo.svg?branch=master)](https://travis-ci.org/Bluetel-Solutions/gomodo)

# Gomodo
The aim of gomodo is to make CLI projects a little easier.

# Basic Usage

```go
package main

import(
	"github.com/Bluetel-Solutions/gomodo"
	"fmt"
)

type Task struct {
	Verbose []bool `short:"v" long:"verbosity" description:"Verbosity flag"`
}

// The Name of the task
func (t *Task) GetName() string {
	return "test-task"
}

// The task Description
func (t *Task) GetDescription() string {
	return "A simple task that does nothing"
}

// The Task Action
func (t *Task) PerformAction(app *gomodo.Application) {
	fmt.Println("Hello World!")
}

func main() {
	app := gomodo.NewApplication("My CLI Name", "v0.1.0")
	app.AddResource(&Task{})

	app.Run()
}

```
