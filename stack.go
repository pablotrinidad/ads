package ads

import "fmt"

// Stack interface
type Stack interface {
	// Top returns the element at the top of the stack.
	Top() (interface{}, error)
	// Size returns the number of elements stored in the stack.
	Size() int
	// Empty removes all elements from the stack.
	Empty()
	// Push a new element into the stack.
	Push(interface{}) error
	// Pop the top element of the stack.
	Pop() (interface{}, error)
}

// ArrayBasedStack is a Stack that uses a fixed-size slice as the
// underlying container.
type ArrayBasedStack struct {
	data      []interface{}
	top, size int
}

// NewArrayBasedStack returns a new Stack of fixed size.
func NewArrayBasedStack(size uint) Stack {
	s := &ArrayBasedStack{}
	s.data = make([]interface{}, size)
	s.top = -1
	s.size = int(size) - 1
	return s
}

// Top returns the element at the top of the stack.
func (s *ArrayBasedStack) Top() (interface{}, error) {
	if s.top < 0 {
		return nil, fmt.Errorf("empty stack")
	}
	return s.data[s.top], nil
}

// Size returns the number of elements stored in the stack.
func (s *ArrayBasedStack) Size() int {
	return s.top + 1
}

// Empty removes all elements from the stack.
func (s *ArrayBasedStack) Empty() {
	s.top = -1
}

// Push a new element into the stack.
func (s *ArrayBasedStack) Push(v interface{}) error {
	if s.top == s.size {
		return fmt.Errorf("stack overflowed, exceeded stack size %d", s.size+1)
	}
	s.top++
	s.data[s.top] = v
	return nil
}

// Pop the top element of the stack.
func (s *ArrayBasedStack) Pop() (interface{}, error) {
	if s.top < 0 {
		return nil, fmt.Errorf("empty stack")
	}
	s.top--
	return s.data[s.top+1], nil
}
