package ads

import "testing"

func logContainerSatisfaction(t *testing.T, ds string, c Container) {
	t.Helper()
	t.Logf("%s satisfies Container interface: %v", ds, c)
}

// TestContainer_Satisfaction verifies (during compilation) that the
// listed data structures satisfy the container interface.
func TestContainer_Satisfaction(t *testing.T) {
	var c Container

	// Dynamic array
	c = NewArray()
	logContainerSatisfaction(t, "Array", c)

	// Doubly Linked list
	ll := NewList()
	logContainerSatisfaction(t, "Doubly Linked List", ll)
}
