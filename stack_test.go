package ads

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func logStackSatisfaction(t *testing.T, ds string, s Stack) {
	t.Helper()
	t.Logf("%s satisfies Stack interface: %v", ds, s)
}

// TestInterfaceSatisfaction verifies (during compilation) that the
// multiple stack implementations satisfy the Stack interface.
func TestInterfaceSatisfaction(t *testing.T) {
	var s Stack

	// Array-based
	s = NewArrayBasedStack(1)
	logStackSatisfaction(t, "ArrayBasedStack", s)
}

// testElementaryMethods will use methods Push, Pop and Top to verify correct implementation.
func testElementaryMethods(t *testing.T, s Stack, n int) {
	t.Helper()
	content, want, got := make([]int, n), make([]int, n), make([]int, 0, n)
	for i := 0; i < n; i++ {
		content[i] = i
		want[i] = n - i - 1
	}
	for _, x := range content {
		if err := s.Push(x); err != nil {
			t.Errorf("Push() produced unexpected error; %v", err)
		}
	}
	for s.Size() > 0 {
		x, err := s.Top()
		if err != nil {
			t.Errorf("Top() produced unexpected error; %v", err)
		}
		y, err := s.Pop()
		if err != nil {
			t.Errorf("Pop() produced unexpected error; %v", err)
		}
		if y != x {
			t.Errorf("Pop():%d != Top():%d", y, x)
		}
		got = append(got, x.(int))
	}
	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("Fill-Flush test failed, got diff %s", diff)
	}
}

func testErrorsOnEmptyList(t *testing.T, s Stack) {
	t.Helper()
	if _, err := s.Top(); err == nil {
		t.Error("Top() returned non-nil error, want error")
	}
	if err := s.Push(1); err == nil {
		t.Error("Push() returned non-nil error, want error")
	}
	if _, err := s.Pop(); err == nil {
		t.Error("Pop() returned non-nil error, want error")
	}
}

func testEmptyProcedure(t *testing.T, s Stack, n int) {
	t.Helper()
	for i := 0; i < n; i++ {
		if err := s.Push(i); err != nil {
			t.Errorf("Push() produced unexpected error; %v", err)
		}
	}
	if s.Size() != n {
		t.Errorf("Failed building stack of size %d, got size %d", n, s.Size())
	}
	s.Empty()
	if s.Size() != 0 {
		t.Errorf("Size(): %d after calling Empty()", s.Size())
	}
	if _, err := s.Pop(); err == nil {
		t.Error("Pop() returned non-nil error, want error")
	}
	if _, err := s.Top(); err == nil {
		t.Error("Top() returned non-nil error, want error")
	}
}

func TestArrayBasedStack_ElementaryOperations(t *testing.T) {
	tests := []int{0, 1, 2, 3, 10, 100}
	for _, n := range tests {
		testName := fmt.Sprintf("size %d", n)
		t.Run(testName, func(t *testing.T) {
			s := NewArrayBasedStack(uint(n))
			testElementaryMethods(t, s, n)
		})
	}
}

func TestArrayBasedStack_Errors(t *testing.T) {
	s := NewArrayBasedStack(0)
	testErrorsOnEmptyList(t, s)
}

func TestArrayBasedStack_Empty(t *testing.T) {
	tests := []int{0, 1, 2, 3, 10, 100}
	for _, n := range tests {
		testName := fmt.Sprintf("size %d", n)
		t.Run(testName, func(t *testing.T) {
			s := NewArrayBasedStack(uint(n))
			testEmptyProcedure(t, s, n)
		})
	}
}
