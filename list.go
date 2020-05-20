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
