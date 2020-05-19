// Package ads exposes the implementation of some algorithms and data structures. This package was
// not created with production purposes.
package ads

// Container represents a specialized data structure that provides access and manipulation methods.
type Container interface {
	// Add a new element to the container.
	Add(interface{})
	// Remove an element from the container.
	Remove(interface{})
	// Contains returns a boolean indicating whether an element is present in the container or not.
	Contains(interface{}) bool
	// Size returns the number of elements stored in the container.
	Size() int
	// Empty removes all elements from the container.
	Empty()
	// Iterator returns a new container iterable.
	Iterator() Iterable
}
