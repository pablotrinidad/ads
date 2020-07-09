package ads

import "fmt"

// Queue interface
type Queue interface {
	// Front element of the queue.
	Front() (interface{}, error)
	// Back element of the queue.
	Back() (interface{}, error)
	// Size returns the number of elements stored in the queue.
	Size() int
	// Empty removes all elements from the queue.
	Empty()
	// Push a new element into the queue.
	Push(interface{}) error
	// Pop the top element of the queue.
	Pop() (interface{}, error)
}

// ArrayBasedQueue is a Queue that uses a fixed-size slice as the underlying container.
type ArrayBasedQueue struct {
	data             []interface{}
	head, tail, size int
}

// NewArrayBasedQueue returns a new Queue of fixed size.
func NewArrayBasedQueue(size uint) Queue {
	q := &ArrayBasedQueue{}
	q.data = make([]interface{}, size+1)
	q.head = 0
	q.tail = 0
	return q
}

func (q *ArrayBasedQueue) isEmpty() bool {
	return q.head == q.tail
}

// movePointer one place to the right, if array limit is met cycles to the start of the array.
func (q *ArrayBasedQueue) movePointer(p int) int {
	return (p + 1) % len(q.data)
}

func (q *ArrayBasedQueue) isFull() bool {
	return q.movePointer(q.tail) == q.head
}

// Front element of the queue.
func (q *ArrayBasedQueue) Front() (interface{}, error) {
	if q.isEmpty() {
		return nil, fmt.Errorf("empty queue")
	}
	return q.data[q.movePointer(q.head)], nil
}

// Back element of the queue.
func (q *ArrayBasedQueue) Back() (interface{}, error) {
	if q.isEmpty() {
		return nil, fmt.Errorf("empty queue")
	}
	return q.data[q.tail], nil
}

// Size returns the number of elements stored in the queue.
func (q *ArrayBasedQueue) Size() int {
	if q.isEmpty() {
		return 0
	}
	if q.tail > q.head {
		return q.tail - q.head
	}
	return len(q.data) - q.head + q.tail
}

// Empty removes all elements from the queue.
func (q *ArrayBasedQueue) Empty() {
	// Avoid memory leaks (free references for garbage collector)
	for i := range q.data {
		q.data[i] = nil
	}
	q.head = 0
	q.tail = 0
}

// Push a new element into the queue.
func (q *ArrayBasedQueue) Push(v interface{}) error {
	if q.isFull() {
		return fmt.Errorf("queue overflowed, exceeded queue size %d", len(q.data)-1)
	}
	q.tail = q.movePointer(q.tail)
	q.data[q.tail] = v
	return nil
}

// Pop the top element of the queue.
func (q *ArrayBasedQueue) Pop() (interface{}, error) {
	if q.isEmpty() {
		return nil, fmt.Errorf("empty queue")
	}
	q.data[q.head] = nil
	q.head = q.movePointer(q.head)
	return q.data[q.head], nil
}
