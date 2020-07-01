package ads

import (
	"fmt"
	"strings"
)

// ListItem holds a value and its next and previous reference.
type ListItem struct {
	Value      interface{}
	prev, next *ListItem
	list       *List
}

// Next returns the next element in the list (if exists).
func (i *ListItem) Next() *ListItem {
	if n := i.next; i.list != nil && n != &i.list.root {
		return n
	}
	return nil
}

// Prev returns the previous element in the list (if exists).
func (i *ListItem) Prev() *ListItem {
	if n := i.prev; i.list != nil && n != &i.list.root {
		return n
	}
	return nil
}

// String returns the item string representation
func (i *ListItem) String() string {
	return fmt.Sprintf("%v", i.Value)
}

// List represents a double linked list.
type List struct {
	// root is used as a sentinel pointer to hold both the head and the tail of the list.
	root   ListItem
	length int
}

// init initializes or clear the linked list.
func (l *List) init() *List {
	l.root.next = &l.root
	l.root.prev = &l.root
	l.length = 0
	return l
}

// initLazy checks if list has to be initialized.
func (l *List) initLazy() {
	if l.root.next == nil {
		l.init()
	}
}

// NewList returns a newly initialized double linked list. This implementation is heavily based in
// built-int containers/list.
func NewList() *List {
	return new(List).init()
}

// Head of the list.
func (l *List) Head() *ListItem {
	if l.length == 0 {
		return nil
	}
	return l.root.next
}

// Tail of the list.
func (l *List) Tail() *ListItem {
	if l.length == 0 {
		return nil
	}
	return l.root.prev
}

// insertAt inserts item `n` after `at`.
func (l *List) insertAt(n, at *ListItem) *ListItem {
	// Let `r` (right) be whatever comes after `at`.
	r := at.next
	r.prev = n
	n.next = r

	// Update `at` to point to `n`
	at.next = n
	n.prev = at

	n.list = l
	l.length++
	return n
}

// Add inserts `v` at the end of the list.
func (l *List) Add(v interface{}) {
	l.initLazy()
	l.insertAt(&ListItem{Value: v}, l.root.prev)
}

// GetItem returns the first occurrence of the given value if exists, else
// returns an error.
func (l *List) GetItem(v interface{}) (*ListItem, error) {
	l.initLazy()
	var item *ListItem
	for n := l.Head(); n != nil; n = n.Next() {
		if n.Value == v {
			item = n
			break
		}
	}
	if item == nil {
		return nil, fmt.Errorf("%v is not present in the list", v)
	}
	return item, nil
}

// RemoveItem deletes given node from list.
func (l *List) RemoveItem(i *ListItem) {
	if i == nil || i.list != l {
		return
	}
	i.prev.next = i.next
	i.next.prev = i.prev
	l.length--

	// Avoid memory leaks (free references for garbage collector)
	i.next = nil
	i.prev = nil
	i.list = nil
}

// Remove deletes all occurrences of from the list.
func (l *List) Remove(v interface{}) {
	l.initLazy()
	n := l.Head()
	for n != nil {
		next := n.Next()
		if n.Value == v {
			l.RemoveItem(n)
		}
		n = next
	}
}

// Empty restart the list status. It does so by deleting each item
// to avoid memory leaks and remove all references.
func (l *List) Empty() {
	n := l.Head()
	for n != nil {
		next := n.Next()
		l.RemoveItem(n)
		n = next
	}
	l.init()
}

func (l *List) Size() int {
	return l.length
}

// Contains returns whether an element is part of the list
func (l *List) Contains(v interface{}) bool {
	for n := l.Head(); n.Next() != nil; n = n.Next() {
		if n.Value == v {
			return true
		}
	}
	return false
}

// String representation of the double linked list.
func (l *List) String() string {
	var b strings.Builder
	for i, n := 0, l.Head(); i < l.length; i, n = i+1, n.Next() {
		b.WriteString(fmt.Sprintf("%v", n.Value))
		if i < l.length-1 {
			b.WriteString(" ↔ ")
		}
	}
	return b.String()
}

func (l *List) Iterator() Iterable {
	return &ListIterable{l: l, n: l.Head()}
}

// ListIterable implements Iterable interface for List
type ListIterable struct {
	n *ListItem
	l *List
}

// Scan returns a boolean indicating if there's a next element or not.
func (i *ListIterable) Scan() bool {
	return i.n != nil && i.n != &i.l.root
}

// Next returns the next element in the iterable.
func (i *ListIterable) Next() (interface{}, error) {
	if !i.Scan() {
		return nil, fmt.Errorf("there isn't a next element")
	}
	v := i.n.Value
	i.n = i.n.Next()
	return v, nil
}
