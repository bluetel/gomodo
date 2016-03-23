package main

// A Resource controls Actions that can be performed
// These are command controllers that contain
// Further actions or a main action.
type Resource interface {
	// Method to get the Resource name for Help method
	GetName() string
}
