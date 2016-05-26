[![Build Status](https://travis-ci.org/Bluetel-Solutions/gomodo.svg?branch=master)](https://travis-ci.org/Bluetel-Solutions/gomodo)

# Gomodo
Gomodo is a GO Lang CLI framework that allows you to express Actions as structs. The purpose of this is to provide you with a basic framework that provides all the bootstrap you need to create large scale command line applications.

# Basic Usage

```go
package main

import(
	"github.com/Bluetel-Solutions/gomodo"
)

type Task struct {
	Verbose []bool `short:"v" long:"verbosity" description:"Verbosity flag"`
}

// The Name of the task
func (t *Task) GetName() string {
	return "test"
}

// The task Description
func (t *Task) GetDescription() string {
	return "A simple task that does nothing"
}

// The Task Action
func (t *Task) PerformAction(app *gomodo.Application) {
	app.Output.WriteLn("Hello World!")
}

func main() {
	app := gomodo.NewApplication("My CLI Name", "My CLI Description", "v0.1.0")
	app.AddResource(&Task{})

	app.Run()
}

```
