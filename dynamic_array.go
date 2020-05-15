// package ads exposes the implementation of some algorithms and data structures. This package was
// not created with production purposes.
package ads

import (
	"fmt"
	"strings"
)

const listDefaultCapacity int = 4

// List implements a dynamic array data structure, just like built-in slices.
type List struct {
	length   int
	capacity int
	data     []interface{}
}

// NewList returns a newly created list of length 0.
func NewList() List {
	return List{}
}

// Add appends a new element to the list.
func (l *List) Add(v interface{}) {
	if l.length == l.capacity {
		l.resize()
	}
	l.data[l.length] = v
	l.length++
}

// Remove an element of the container (if exists).
func (l *List) Remove(v interface{}) {
	for i := l.length - 1; i >= 0; i-- {
		if l.data[i] == v {
			l.RemoveIth(i)
		}
	}
}

// RemoveIth removes the ith element of the array.
func (l *List) RemoveIth(i int) error {
	if !l.validIndex(i) {
		return fmt.Errorf("index %d out of range", i)
	}
	copy(l.data[i:], l.data[i+1:l.length])
	l.data[l.length-1] = nil
	l.length--
	return nil
}

// Get returns the ith-element of the array.
func (l *List) Get(i int) (interface{}, error) {
	if !l.validIndex(i) {
		return nil, fmt.Errorf("index %d out of range", i)
	}
	return l.data[i], nil
}

// Contains returns whether an element is in the container or not. Implementation is linear search
// since no assumptions can be made about the order of the data.
func (l *List) Contains(v interface{}) bool {
	for i := 0; i < l.length; i++ {
		if l.data[i] == v {
			return true
		}
	}
	return false
}

// Iterator returns an array iterator
func (l *List) Iterator() Iterable {
	return ListIterable{i: 0, l: l}
}

// Size returns the length of the array.
func (l *List) Size() int {
	return int(l.length)
}

// Empty removes all the elements in the array.
func (l *List) Empty() {
	l.length = 0
	l.capacity = 0
	l.data = nil
}

// Stringer returns a string representation of the array content.
func (l List) String() string {
	var b strings.Builder
	b.WriteString("[")
	for i := 0; i < l.length; i++ {
		b.WriteString(fmt.Sprintf("%v", l.data[i]))
		if i < l.length-1 {
			b.WriteString(", ")
		}
	}
	b.WriteString("]")
	return b.String()
}

func (l *List) validIndex(i int) bool {
	return i >= 0 && i < l.length
}

// resize internal array by a factor of ~1.125. See CPython's list_resize method for reference
// https://github.com/python/cpython/blob/master/Objects/listobject.c#L36.
func (l *List) resize() {
	switch l.capacity {
	case 0:
		l.capacity = listDefaultCapacity
	default:
		l.capacity = (l.capacity + (l.capacity >> 3) + 6) &^ 3
	}
	newData := make([]interface{}, l.capacity)
	copy(newData, l.data)
	l.data = newData
}

// ListIterable implements Iterable interface for List.
type ListIterable struct {
	i int
	l *List
}

// Scan returns a boolean indicating if there's a next element or not.
func (i ListIterable) Scan() bool {
	return i.i < i.l.length
}

// Next returns the next element in the iterable.
func (i ListIterable) Next() (interface{}, error) {
	return i.l.Get(i.i)
}
