package ads

import (
	"fmt"
	"strings"
)

const arrayDefaultCapacity int = 4

// Array implements a dynamic array data structure, just like built-in slices.
type Array struct {
	length   int
	capacity int
	data     []interface{}
}

// NewArray returns a newly created array of length 0.
func NewArray() *Array {
	return &Array{}
}

// Add appends a new element to the init.
func (a *Array) Add(v interface{}) {
	if a.length == a.capacity {
		a.resize()
	}
	a.data[a.length] = v
	a.length++
}

// Remove an element of the array (if exists).
func (a *Array) Remove(v interface{}) {
	for i := a.length - 1; i >= 0; i-- {
		if a.data[i] == v {
			a.RemoveIth(i)
		}
	}
}

// RemoveIth removes the ith element of the array.
func (a *Array) RemoveIth(i int) error {
	if !a.validIndex(i) {
		return fmt.Errorf("index %d out of range", i)
	}
	copy(a.data[i:], a.data[i+1:a.length])
	a.data[a.length-1] = nil
	a.length--
	return nil
}

// Get returns the ith-element of the array.
func (a *Array) Get(i int) (interface{}, error) {
	if !a.validIndex(i) {
		return nil, fmt.Errorf("index %d out of range", i)
	}
	return a.data[i], nil
}

// Contains returns whether an element is in the array or not. Implementation is linear search
// since no assumptions can be made about the order of the data.
func (a *Array) Contains(v interface{}) bool {
	for i := 0; i < a.length; i++ {
		if a.data[i] == v {
			return true
		}
	}
	return false
}

// Size returns the length of the array.
func (a *Array) Size() int {
	return int(a.length)
}

// Empty removes all the elements in the array.
func (a *Array) Empty() {
	a.length = 0
	a.capacity = 0
	a.data = nil
}

// Stringer returns a string representation of the array content.
func (a Array) String() string {
	var b strings.Builder
	b.WriteString("[")
	for i := 0; i < a.length; i++ {
		b.WriteString(fmt.Sprintf("%v", a.data[i]))
		if i < a.length-1 {
			b.WriteString(", ")
		}
	}
	b.WriteString("]")
	return b.String()
}

func (a *Array) validIndex(i int) bool {
	return i >= 0 && i < a.length
}

// resize internal array by a factor of ~1.125. See CPython's list_resize method for reference
// https://github.com/python/cpython/blob/master/Objects/listobject.c#L36.
func (a *Array) resize() {
	switch a.capacity {
	case 0:
		a.capacity = arrayDefaultCapacity
	default:
		a.capacity = (a.capacity + (a.capacity >> 3) + 6) &^ 3
	}
	newData := make([]interface{}, a.capacity)
	copy(newData, a.data)
	a.data = newData
}

// Iterator returns an array iterator
func (a *Array) Iterator() Iterable {
	return &ArrayIterable{i: 0, a: a}
}

// ArrayIterable implements Iterable interface for Array.
type ArrayIterable struct {
	i int
	a *Array
}

// Scan returns a boolean indicating if there's a next element or not.
func (i *ArrayIterable) Scan() bool {
	return i.i < i.a.length
}

// Next returns the next element in the iterable.
func (i *ArrayIterable) Next() (interface{}, error) {
	v, err := i.a.Get(i.i)
	if err != nil {
		return nil, err
	}
	i.i++
	return v, nil
}
